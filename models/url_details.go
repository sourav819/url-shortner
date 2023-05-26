package models

import (
	"time"

	"gorm.io/gorm"
)

type UrlInfo struct {
	ID          uint64          `json:"id" gorm:"primarykey"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
	OriginalUrl *string         `json:"original_url"`
	ShortUrl    *string         `json:"short_url"`
	Code        *string         `json:"code"`
}

type UrlInfoRepo struct {
	DB *gorm.DB
}

func (u *UrlInfoRepo) Create(ul *UrlInfo) error {
	return u.CreateWithTx(u.DB, ul)
}

func (u *UrlInfoRepo) CreateWithTx(tx *gorm.DB, ul *UrlInfo) error {
	err := tx.Create(ul).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlInfoRepo) Update(tx *gorm.DB,ul *UrlInfo, ID uint64) error {
	return u.UpdateWithTx(tx, ul, ID)
}

func (u *UrlInfoRepo) UpdateWithTx(tx *gorm.DB, ul *UrlInfo, ID uint64) error {
	err := tx.Model(&UrlInfo{ID: ID}).Updates(ul).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlInfoRepo) GetById(code string) (*UrlInfo, int64, error) {
	return u.GetWithTx(u.DB, &UrlInfo{Code: &code})
}

func (u *UrlInfoRepo) GetUrlById(url string) (*UrlInfo, error) {
	return u.GetUrlDetails(u.DB, url)
}

func (u *UrlInfoRepo) GetWithTx(tx *gorm.DB, code *UrlInfo) (*UrlInfo, int64, error) {
	var cd UrlInfo
	fetchedRecords := tx.Model(&UrlInfo{}).Where(code).Find(&cd)
	// if err != nil {
	// 	return nil, err
	// }
	return &cd, fetchedRecords.RowsAffected, fetchedRecords.Error
}

func (u *UrlInfoRepo) GetUrlDetails(tx *gorm.DB, url string) (*UrlInfo, error) {
	var o UrlInfo
	err := tx.Model(&UrlInfo{}).Where(&UrlInfo{OriginalUrl: &url}).First(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}
