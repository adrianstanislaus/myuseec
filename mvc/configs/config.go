package configs

import (
	"myuseek/models/albums"
	"myuseek/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:KucingSQL420@tcp(127.0.0.1:3306)/myuseek?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB failed connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(&albums.Album{}, &users.User{})
}
