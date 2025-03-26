package api

import (
	"taskape-rest-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handlers.Handlers) {
	app.Get("/ping", h.Auth.Ping)
	app.Post("/sendVerificationCode", h.Auth.SendVerificationCode)
	app.Post("/checkVerificationCode", h.Auth.CheckVerificationCode)
	app.Post("/validateToken", h.Auth.ValidateToken)
	app.Post("/refreshToken", h.Auth.RefreshToken)

	app.Get("/users/:userID", h.User.GetUser)
	app.Post("/checkHandleAvailability", h.User.CheckHandleAvailability)
	app.Post("/registerNewProfile", h.User.RegisterNewProfile)

	app.Post("/submitTask", h.Task.SubmitTask)
	app.Post("/submitTasksBatch", h.Task.SubmitTasksBatch)
	app.Get("/users/:userID/tasks", h.Task.GetUserTasks)
	app.Post("/updateTask", h.Task.UpdateTask)
	app.Post("/updateTaskOrder", h.Task.UpdateTaskOrder)

	app.Post("/searchUsers", h.Friend.SearchUsers)
	app.Post("/sendFriendRequest", h.Friend.SendFriendRequest)
	app.Post("/respondToFriendRequest", h.Friend.RespondToFriendRequest)
	app.Get("/users/:userID/friends", h.Friend.GetUserFriends)
	app.Get("/users/:userID/friendRequests", h.Friend.GetFriendRequests)
}
