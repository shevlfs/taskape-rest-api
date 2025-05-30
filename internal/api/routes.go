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
	app.Post("/confirmTaskCompletion", h.Task.ConfirmTaskCompletion)

	app.Post("/searchUsers", h.Friend.SearchUsers)
	app.Post("/sendFriendRequest", h.Friend.SendFriendRequest)
	app.Post("/respondToFriendRequest", h.Friend.RespondToFriendRequest)
	app.Get("/users/:userID/friends", h.Friend.GetUserFriends)
	app.Get("/users/:userID/friendRequests", h.Friend.GetFriendRequests)

	app.Get("/users/:userID/events", h.Event.GetUserEvents)

	app.Post("/events/:eventID/like", h.Event.LikeEvent)
	app.Delete("/events/:eventID/like", h.Event.UnlikeEvent)
	app.Get("/events/:eventID/comments", h.Event.GetEventComments)
	app.Post("/events/:eventID/comments", h.Event.AddEventComment)
	app.Delete("/events/:eventID/comments/:commentID", h.Event.DeleteEventComment)

	app.Post("/getUsersBatch", h.User.GetUsersBatch)
	app.Post("/getUsersTasksBatch", h.Task.GetUsersTasksBatch)
	app.Post("/editUserProfile", h.User.EditUserProfile)

	app.Get("/users/:userID/relatedEvents", h.Event.GetUserRelatedEvents)

	app.Post("/createGroup", h.User.CreateGroup)
	app.Get("/groups/:groupID/tasks", h.User.GetGroupTasks)
	app.Post("/inviteToGroup", h.User.InviteToGroup)
	app.Post("/acceptGroupInvite", h.User.AcceptGroupInvite)
	app.Post("/kickUserFromGroup", h.User.KickUserFromGroup)
	app.Get("/users/:userID/groups", h.User.GetUserGroups)
	app.Get("/users/:userID/groupInvitations", h.User.GetGroupInvitations)

	app.Get("/users/:userID/streak", h.User.GetUserStreak)
}
