package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("shivam:shivam123@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"),&gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
