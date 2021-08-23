package webservice

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeroidentidad/gopherbot/webservice/challenges"
	"github.com/zeroidentidad/gopherbot/webservice/storage"
)

func Start(port string) {
	app := createApp()
	app.Listen(":" + port)
}

func createApp() *fiber.App {
	app := fiber.New()
	db := storage.Get()
	challengeSvc := challenges.ChallengeService{
		ChallengeGateway: challenges.ChallengeGateway{
			DB: db,
		},
	}

	// TODO: build frontend
	app.Static("/", "./webservice/frontend")
	app.Add(http.MethodGet, "/health", HealthCheckHandler)
	app.Add(http.MethodPost, "/challenges", challengeSvc.CreateChallengeHandler)
	return app
}
