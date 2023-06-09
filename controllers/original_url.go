package controllers

import (
	"net/http"
	"sync"
	"url-shortner/constants"
	"url-shortner/models"
	"url-shortner/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExtractInfo struct {
	Code        string `json:"code"`
	UrlCode     uint64 `json:"url_code"`
	OriginalUrl string `json:"original_url"`
}

var cache = &sync.Map{}

func (b *BaseController) GetOriginalUrl(c *gin.Context) {
	var (
		code        string
		urlRepo     = models.InitUrlDetailsRepo(b.DB)
		errResponse = constants.ErrorEntity{}
	)
	//fetch
	code = c.Param("code")

	if code == "" {
		logger.Error("code is mandatory in short url ")
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	longURL, ok := cache.Load(code)
	if ok {
		logger.Info("using value stored in cache...")
		c.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}

	codeInfo, rowsEffected, err := urlRepo.GetById(code)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("unable to fetch code for the original url ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	if rowsEffected == 0 {
		logger.Error("code not found in records ")
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	cacheStore := *codeInfo.OriginalUrl
	cache.Store(code, cacheStore)
	OriginalUrl := *codeInfo.OriginalUrl
	redirectUrl := OriginalUrl[0 : len(OriginalUrl)-6]
	redirectUrl += "v1/" + OriginalUrl[len(OriginalUrl)-6:]
	c.Redirect(http.StatusMovedPermanently, redirectUrl)

}
