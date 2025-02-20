package model

import "time"

type MstUser struct {
	Id              string         `gorm:"size:36;primaryKey"`
	FirstName       string         `gorm:"size:100;not null"`
	LastName        string         `gorm:"size:100;"`
	Email           string         `gorm:"size:255;not null;unique"`
	Phone           string         `gorm:"size:100;not null;unique"`
	Password        string         `gorm:"size:255;not null"`
	Address         []MstAddresses `gorm:"foreignKey:UserId"`
	EmailVeriviedAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type MstAddresses struct {
	Id          string `gorm:"size:36;primaryKey"`
	UserId      string `gorm:"size:36;index;not null"`
	Addresses1  string `gorm:"size:255;not null"`
	Addresses2  string `gorm:"size:255;"`
	CityId      string `gorm:"size:50;not null"`
	ProvincesId string `gorm:"size:50;not null"`
	PostCode    string `dorm:"size:50;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
