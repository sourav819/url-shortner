package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64          `json:"id" gorm:"primarykey"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *time.Time      `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	UUID      string          `json:"uuid" gorm:"unique"`

	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type userRepo struct {
	DB *gorm.DB
}

func (u *userRepo) GetUserDetails(email, phoneNum string) (int64, error) {
	return u.GetwithTx(u.DB, &User{Email: email, PhoneNumber: phoneNum})
}

func (u *userRepo) GetwithTx(tx *gorm.DB, email *User) (int64, error) {
	var ud User
	result := tx.Model(&User{}).Where(email).First(&ud)
	return result.RowsAffected, result.Error
}

// func (u *UserDetailsRepo) CheckUserBasedOnEmail(email string) (int64, error) {
// 	return u.GetwithTx(u.DB, &User{Email: &email})
// }

// func (u *UserDetailsRepo) CheckUserBasedOnPhone(phoneNum string) (int64, error) {
// 	return u.GetwithTx(u.DB, &User{PhoneNumber: &phoneNum})
// }

func (u *userRepo) CheckDetailsForLogin(tx *gorm.DB, email, phoneNum string) (*User, error) {
	var us User
	err := tx.Model(&User{}).Where(`email=?`, email).Or(`phone_number=?`, phoneNum).First(&us).Error
	return &us, err
}
