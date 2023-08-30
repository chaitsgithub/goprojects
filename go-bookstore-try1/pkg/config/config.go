package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func DBConnect() {
	d, err := gorm.Open("mysql", "root:admin@/bookstoredb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetConnection() *gorm.DB {
	return db
}
