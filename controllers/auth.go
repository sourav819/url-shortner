package controllers

import (
	"net/http"
	"url-shortner/auth"
	"url-shortner/constants"
	"url-shortner/models"
	"url-shortner/pkg/logger"
	"url-shortner/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (b *BaseController) SignUpUser(c *gin.Context) {

	var (
		request     = SignUpRequest{}
		response    = AuthResponse{}
		err         error
		errMsg      = constants.ErrorEntity{}
		userRepo    = models.InitUserRepo(b.DB)
		userDetails = models.User{}
	)

	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest,
			errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	tx := b.DB.Begin()
	recordCount, err := userRepo.GetUserDetails(request.Email, request.PhoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("unable to get user details")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	if recordCount != 0 {
		logger.Error("user already exists")
		c.AbortWithStatusJSON(http.StatusForbidden,
			errMsg.GenerateError(http.StatusForbidden, "user already exists"))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("unable to hash password")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	randomId, err := utils.CreateNanoID(20)
	if err != nil {
		logger.Error("unable to generate random uuid")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}

	passwd := string(hashedPassword)
	userDetails.FirstName = request.FirstName
	userDetails.LastName = request.LastName
	userDetails.Email = request.Email
	userDetails.PhoneNumber = request.PhoneNumber
	userDetails.Password = passwd
	userDetails.UUID = "usr_" + randomId

	err = tx.Create(&userDetails).Error
	if err != nil {
		logger.Error("unable to insert record of the user")
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))

	}
	err = tx.Commit().Error
	if err != nil {
		b.Log.Error("unable to commit the code and url ", err)
		tx.Rollback()
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	token, err := auth.GenerateToken(b.Config.JWTConfig.JWTTokenValidTimeInHour, userDetails.Email, userDetails.FirstName, userDetails.LastName, userDetails.UUID, b.Config.JWTConfig.JWTSecret)
	if err != nil {
		logger.Error("unable to generate token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	response.Message = "User signed up successfully"
	response.Token = token
	c.JSON(http.StatusCreated, response)
}

func (b *BaseController) Login(c *gin.Context) {
	var (
		request  = LoginRequest{}
		response = AuthResponse{}
		err      error
		errMsg   = constants.ErrorEntity{}
		userRepo = models.InitUserRepo(b.DB)
	)
	err = c.ShouldBindJSON(&request)
	if err != nil {
		logger.Error("invalid request")
		c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "invalid request"))
		return
	}
	tx := b.DB.Begin()
	//checking if the phone number or email exists
	usrDetails, err := userRepo.CheckDetailsForLogin(tx, request.Email, request.PhoneNumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Info("user doesnot exists")
			c.AbortWithStatusJSON(http.StatusBadRequest, errMsg.GenerateError(http.StatusBadRequest, "emailId or phone number does not exists, please signup"))
			return
		} else {
			logger.Error("unable to fetch details for user")
			c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
			return
		}
	}
	//check password
	if !utils.VerifyPassword(usrDetails.Password, request.Password) {
		logger.Error("password doesnot match")
		c.AbortWithStatusJSON(http.StatusForbidden, errMsg.GenerateError(http.StatusForbidden, "incorrect password"))
		return
	}
	token, err := auth.GenerateToken(b.Config.JWTConfig.JWTTokenValidTimeInHour, usrDetails.Email, usrDetails.FirstName, usrDetails.LastName, usrDetails.UUID, b.Config.JWTConfig.JWTSecret)
	if err != nil {
		logger.Error("unable to generate token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, errMsg.GenerateError(http.StatusInternalServerError, "something went wrong"))
		return
	}
	response.Message = "logged in successfully"
	response.Token = token
	c.JSON(http.StatusOK, response)

}
