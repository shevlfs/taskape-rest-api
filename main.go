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
	app.Post("/submitTask", handler.SubmitTask)
	app.Post("/submitTasksBatch", handler.SubmitTasksBatch)

	app.Post("/submitTask", handler.SubmitTask)
	app.Post("/submitTasksBatch", handler.SubmitTasksBatch)
	app.Get("/users/:userID/tasks", handler.GetUserTasks)
	app.Get("/users/:userID", handler.GetUser)

	app.Post("/checkHandleAvailability", handler.CheckHandleAvailability)

	app.Post("/updateTask", handler.UpdateTask)

	app.Post("/updateTaskOrder", handler.UpdateTaskOrder)

	log.Fatal(app.Listen(":8080"))
}
