package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Get() *gorm.DB {
	if DB == nil {
		DB = get()
	}
	return DB
}

func get() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
