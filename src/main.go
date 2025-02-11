package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	app.Get("/ping", ping)
	log.Fatal(app.Listen(":8080"))
}
