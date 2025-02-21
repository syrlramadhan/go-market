package controller

import (
	"github.com/syrlramadhan/go-market/app/model"
	"gorm.io/gorm"
)

func GetProductBySlug(slug string, db *gorm.DB) (model.MstProducts, error) {
	var product model.MstProducts
	result := db.Preload("ProductImages").Where("slug = ?", slug).First(&product)

	if result.Error != nil {
		return model.MstProducts{}, result.Error
	}

	return product, nil
}

func GetProduct(db *gorm.DB) ([]model.MstProducts, error) {
	var product []model.MstProducts
	result := db.Preload("ProductImages").Find(&product).Limit(15)

	if result.Error != nil {
		return []model.MstProducts{}, result.Error
	}

	return product, nil
}