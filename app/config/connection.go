package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "root@tcp(localhost:3306)/go_market_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	return db
}
