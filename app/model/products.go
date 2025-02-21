package model

import "time"

type MstProducts struct {
	Id              string             `gorm:"size:36;primaryKey"`
	Name            string             `gorm:"size:255;not null"`
	Slug            string             `gorm:"size:255;not null;"`
	Price           float64            `gorm:"not null"`
	Stock           int64              `gorm:"size:50;not null"`
	Weight          float64            `gorm:"not null"`
	SortDescription string             `gorm:"size:100;not null"`
	Description     string             `gorm:"type:text;not null"`
	ProductImages   []MstProductImages `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MstProductImages struct {
	Id        string `gorm:"size:36;primaryKey"`
	ProductId string `gorm:"size:36;index;not null"`
	Path      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
