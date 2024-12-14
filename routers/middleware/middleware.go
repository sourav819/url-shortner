package middleware

import (
	"url-shortner/auth"
	"url-shortner/constants"
	"url-shortner/models"
	"url-shortner/pkg/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authentication(JWTSecret string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			errMsg = constants.ErrorEntity{}
			users  models.User
		)
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			logger.Error("token not present")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg.GenerateError(http.StatusUnauthorized, "auth header cannot be empty"))
			return
		}
		AuthorizationKey := strings.Split(tokenHeader, " ")
		tokenDetails, err := auth.ValidateToken(AuthorizationKey[1], JWTSecret)
		if err != nil {
			logger.Error("invalid token received")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg.GenerateError(http.StatusUnauthorized, "Please Login "))
			return
		}

		if db.Model(&models.User{}).Where(&models.User{UUID: tokenDetails.UUID}).Find(&users).RowsAffected == 0 {

			logger.Error("user does not exists")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg.GenerateError(http.StatusUnauthorized, "Please Signup"))
			return
		}

		c.Set("user", users)
		c.Next()
	}
}
