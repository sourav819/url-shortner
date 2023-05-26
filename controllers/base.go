package controllers

import (
	"url-shortner/pkg/config"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type BaseController struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config config.Config
}
