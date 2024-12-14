package routers

import (
	"url-shortner/controllers"
	"url-shortner/pkg/config"
	"url-shortner/routers/middleware"

	"github.com/sirupsen/logrus"
)

func v1Routes(app config.AppConfig) {
	ctrl := controllers.BaseController{
		Config: app.Config,
		DB:     app.DB,
		Log:    logrus.New(),
	}
	v1 := app.Router.Group("/v1")

	//urls entity
	v1.GET("/:code", ctrl.GetOriginalUrl)
	usersGroup := v1.Group("/users")
	usersGroup.POST("/signup", ctrl.SignUpUser)
	usersGroup.POST("/signin", ctrl.Login)

	urlGroup := v1.Group("url", middleware.Authentication(ctrl.Config.JWTConfig.JWTSecret, ctrl.DB))
	urlGroup.POST("/generate", ctrl.GenerateShortUrl)

}
