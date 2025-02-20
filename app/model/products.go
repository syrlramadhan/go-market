package model

import "time"

type MstProducts struct {
	Id              string  `gorm:"size:36;primaryKey"`
	Name            string  `gorm:"size:255;not null"`
	Price           float64 `gorm:"not null"`
	Stock           int64   `gorm:"size:50;not null"`
	Weight          float64 `gorm:"not null"`
	SortDescription string  `gorm:"size:100;not null"`
	Description     string  `gorm:"type:text;not null"`
	ImagePath       string  `gorm:"size:255;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
