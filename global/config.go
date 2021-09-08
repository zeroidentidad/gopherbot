package global

import "os"

var (
	TOKEN   = os.Getenv("TOKEN")
	PORT    = os.Getenv("PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_NAME = os.Getenv("DB_NAME")
)

const (
	DB_HOST   string = "mysql-zeroidentidad.alwaysdata.net"
	port_conf string = "3000"
)

func Port() string {
	if PORT == "" {
		PORT = port_conf
	}
	return PORT
}
