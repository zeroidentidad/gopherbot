package config

import "os"

const (
	Token    string = "NzY3OTc5MDk4NzY4NDA4NTg2.X45yRQ.U-ue6pQabGEFui8pFqIqtFvSJ94"
	portConf string = "3000"
)

var PORT string = os.Getenv("PORT")

func Port() string {
	if PORT == "" {
		PORT = portConf
	}
	return PORT
}
