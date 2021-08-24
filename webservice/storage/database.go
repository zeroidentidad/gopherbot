package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/zeroidentidad/gopherbot/global"
	"gorm.io/driver/mysql"
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
	dsn := fmt.Sprintf(global.DB_USER+":"+global.DB_PASS+"@tcp(%s:3306)/%s?collation=utf8mb4_general_ci&parseTime=True&loc=Local", global.DB_HOST, global.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// pool settings:
	pool, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	return db
}
