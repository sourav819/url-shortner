package controllers

import (
	"net/http"
	"net/url"
	"url-shortner/constants"
	"url-shortner/models"
	"url-shortner/objects"
	"url-shortner/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (b *BaseController) GenerateShortUrl(c *gin.Context) {

	var (
		request     objects.GenerateURL
		err         error
		urlRepo     = models.InitUrlDetailsRepo(b.DB)
		urlDetails  = models.UrlInfo{}
		baseCode    string
		errResponse = constants.ErrorEntity{}
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		b.Log.Error("unable to bind json ", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse.GenerateError(http.StatusBadRequest, "invalid request"))
		return

	}
	_, urlErr := url.ParseRequestURI(request.OriginalUrl)
	if urlErr != nil {
		b.Log.Error("invalid url ", urlErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse.GenerateError(http.StatusBadRequest, "invalid url provided"))
		return
	}
	//check if the url is already present in db
	tx := b.DB.Begin()
	existUrl, err := urlRepo.GetUrlById(request.OriginalUrl)
	if err != nil && err != gorm.ErrRecordNotFound {
		b.Log.Error("unable to fetch record for a url ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	if existUrl != nil {
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"message":   "short url exists",
			"short-url": existUrl.ShortUrl,
		})
		return
	}

	//insert long url in table first
	urlDetails.OriginalUrl = &request.OriginalUrl
	err = tx.Create(&urlDetails).Error
	if err != nil {
		b.Log.Error("unable to create record for url ", err)
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	if request.CustomShort != "" {
		baseCode = request.CustomShort
	} else {
		baseCode, err = utils.CreateNanoID(12)
		if err != nil {
			b.Log.Error("unable to create code for short url ", err)
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
		baseCode = baseCode[:6]
	}

	//check if the code exists in database
	_, rowsEffected, err := urlRepo.GetById(baseCode)
	if err != nil && err != gorm.ErrRecordNotFound {
		b.Log.Error("unable to fetch code for the original url ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	if rowsEffected != 0 {
		b.Log.Error("custom code already exists  ")
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusForbidden, errResponse.GenerateError(http.StatusForbidden, "action not allowed"))
		return
	}

	finalShortUrl := b.Config.URLDetails.BaseUrl + baseCode
	urlDetails.ShortUrl = &finalShortUrl
	urlDetails.Code = &baseCode
	err = urlRepo.Update(tx, &urlDetails, urlDetails.ID)
	if err != nil {
		b.Log.Error("unable to update record for url ", err)
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	err = tx.Commit().Error
	if err != nil {
		b.Log.Error("unable to commit the code and url ", err)
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":   "short url generated successfully",
		"short-url": finalShortUrl,
	})
}
