package models

import "gorm.io/gorm"

type IUrlDetails interface {
	Create(ul *UrlInfo) error
	CreateWithTx(tx *gorm.DB, ul *UrlInfo) error
	Update(tx *gorm.DB, ul *UrlInfo, ID uint64) error
	UpdateWithTx(tx *gorm.DB, ul *UrlInfo, ID uint64) error
	GetUrlDetails(tx *gorm.DB, url string) (*UrlInfo, error)
	GetById(code string) (*UrlInfo, int64, error)
	GetUrlById(url string) (*UrlInfo, error)
	GetWithTx(tx *gorm.DB, code *UrlInfo) (*UrlInfo, int64, error)
}

type IUser interface {
	GetUserDetails(email, phoneNum string) (int64, error)
	GetwithTx(tx *gorm.DB, email *User) (int64, error)
	CheckDetailsForLogin(tx *gorm.DB, email, phoneNum string) (*User, error)
}
