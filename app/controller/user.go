package controller

import (
	"github.com/syrlramadhan/go-market/app/model"
	"gorm.io/gorm"
)

func CreateUser(user model.MstUser, db *gorm.DB) error {
	query := db.Select("Id", "FirstName", "LastName", "Email", "Password").Create(&user)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func GetUserByEmail(email string, db *gorm.DB) (model.MstUser, error) {
	var user model.MstUser
	query := db.Where("email = ?", email).First(&user)
	if query.Error != nil {
		return model.MstUser{}, query.Error
	}

	return user, nil
}