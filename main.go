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
	app.Post("/validateToken", handler.VerifyUserToken)
	app.Post("/refreshToken", handler.RefreshToken)
	app.Post("/sendVerificationCode", handler.VerificationCodeRequestRoute)
	app.Post("/checkVerificationCode", handler.CheckVerificationCode)
	app.Post("/registerNewProfile", handler.RegisterNewProfile)
	log.Fatal(app.Listen(":8080"))
}
