package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("test.db"))

	if err != nil {
		panic("Connection to DB failed!")
	}

	database.AutoMigrate(&User{})
	DB = database
}
