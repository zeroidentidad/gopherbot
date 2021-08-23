package storage

import (
	"log"

	"github.com/zeroidentidad/gopherbot/webservice/challenges"
)

func Migrate() {
	db := Get()

	err := db.AutoMigrate(challenges.Challenge{})
	if err != nil {
		log.Fatal("Cannot do migration:", err.Error())
	}
}
