package models

import "gorm.io/gorm"

func InitUrlDetailsRepo(DB *gorm.DB) IUrlDetails {
	return &UrlInfoRepo{
		DB: DB,
	}
}
