package controllers

import (
	"url-shortner/pkg/config"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Config config.Config
	Validator             *validator.Validate
	Translator            *ut.Translator
}
