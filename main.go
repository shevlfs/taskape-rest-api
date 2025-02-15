package main

import (
	"log"

	"taskape-rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	handler := routes.New()
	app.Get("/ping", handler.Ping)
	app.Post("/sendVerificationCode", handler.VerificationCodeRequestRoute)
	app.Post("/checkVerificationCode", handler.CheckVerificationCode)
	log.Fatal(app.Listen(":8080"))
}
