package handlers

import (
	"taskape-rest-api/internal/config"
	proto "taskape-rest-api/proto"
)

type Handlers struct {
	User   *UserHandler
	Task   *TaskHandler
	Auth   *AuthHandler
	Friend *FriendHandler
	Event *EventHandler
}

func NewHandlers(client proto.BackendRequestsClient, cfg *config.Config) *Handlers {
	return &Handlers{
		User:   &UserHandler{BackendClient: client},
		Task:   &TaskHandler{BackendClient: client},
		Auth:   NewAuthHandler(client, cfg),
		Friend: &FriendHandler{BackendClient: client},
	}
}
