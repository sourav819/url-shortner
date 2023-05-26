package routers

import (
	"url-shortner/controllers"
	"url-shortner/pkg/config"

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
	v1.POST("/generate", ctrl.GenerateShortUrl)
	v1.GET("/:code", ctrl.GetOriginalUrl)
}
