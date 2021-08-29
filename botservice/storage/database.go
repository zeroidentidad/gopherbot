package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeroidentidad/gopherbot/global"
)

var db *sql.DB

func MySqlConn() *sql.DB {
	if db != nil {
		return db
	}

	dsn := fmt.Sprintf(global.DB_USER+":"+global.DB_PASS+"@tcp(%s:3306)/%s?collation=utf8mb4_general_ci&parseTime=true", global.DB_HOST, global.DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
