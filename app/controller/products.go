package controller

import (
	"github.com/syrlramadhan/go-market/app/model"
	"gorm.io/gorm"
)

func GetProductByID(id string, db *gorm.DB) (model.MstProducts, error) {
	var product model.MstProducts
	result := db.Where("id = ?", id).First(&product)

	if result.Error != nil {
		return model.MstProducts{}, result.Error
	}

	return product, nil
}

func GetProduct(db *gorm.DB) ([]model.MstProducts, error) {
	var product []model.MstProducts
	result := db.Find(&product)

	if result.Error != nil {
		return []model.MstProducts{}, result.Error
	}

	return product, nil
}