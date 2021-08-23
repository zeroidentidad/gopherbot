package webservice

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(ctx *fiber.Ctx) error {
	_, err := ctx.WriteString("Service is ok")
	return err
}
