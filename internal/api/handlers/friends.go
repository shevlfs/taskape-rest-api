package handlers

import (
	"context"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

type FriendHandler struct {
	BackendClient proto.BackendRequestsClient
}

func (h *FriendHandler) SearchUsers(c *fiber.Ctx) error {
	var request dto.SearchUsersRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SearchUsersResponse{
			Success: false,
			Users:   nil,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.Query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SearchUsersResponse{
			Success: false,
			Users:   nil,
			Message: "Search query is required",
		})
	}

	if request.Limit <= 0 {
		request.Limit = 10
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.SearchUsers(ctx, &proto.SearchUsersRequest{
		Query: request.Query,
		Limit: int32(request.Limit),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.SearchUsersResponse{
			Success: false,
			Users:   nil,
			Message: "Failed to search users: " + err.Error(),
		})
	}

	users := make([]dto.UserSearchResult, len(resp.Users))
	for i, user := range resp.Users {
		users[i] = dto.UserSearchResult{
			ID:             user.Id,
			Handle:         user.Handle,
			ProfilePicture: user.ProfilePicture,
			Color:          user.Color,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.SearchUsersResponse{
		Success: true,
		Users:   users,
		Message: "",
	})
}

func (h *FriendHandler) SendFriendRequest(c *fiber.Ctx) error {
	var request dto.SendFriendRequestRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.SenderID == "" || request.ReceiverID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Sender ID and Receiver ID are required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.SendFriendRequest(ctx, &proto.SendFriendRequestRequest{
		SenderId:   request.SenderID,
		ReceiverId: request.ReceiverID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Failed to send friend request: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.SendFriendRequestResponse{
		Success:   resp.Success,
		RequestID: resp.RequestId,
		Message:   resp.Error,
	})
}

func (h *FriendHandler) RespondToFriendRequest(c *fiber.Ctx) error {
	var request dto.RespondToFriendRequestRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.RequestID == "" || request.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Request ID and User ID are required",
		})
	}

	if request.Response != "accept" && request.Response != "reject" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Response must be 'accept' or 'reject'",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.RespondToFriendRequest(ctx, &proto.RespondToFriendRequestRequest{
		RequestId: request.RequestID,
		UserId:    request.UserID,
		Response:  request.Response,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Failed to respond to friend request: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.RespondToFriendRequestResponse{
		Success: resp.Success,
		Message: resp.Error,
	})
}

func (h *FriendHandler) GetUserFriends(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUserFriendsResponse{
			Success: false,
			Friends: nil,
			Message: "User ID is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.GetUserFriendsResponse{
			Success: false,
			Friends: nil,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUserFriends(ctx, &proto.GetUserFriendsRequest{
		UserId: userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetUserFriendsResponse{
			Success: false,
			Friends: nil,
			Message: "Failed to get user friends: " + err.Error(),
		})
	}

	friends := make([]dto.Friend, len(resp.Friends))
	for i, friend := range resp.Friends {
		friends[i] = dto.Friend{
			ID:             friend.Id,
			Handle:         friend.Handle,
			ProfilePicture: friend.ProfilePicture,
			Color:          friend.Color,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetUserFriendsResponse{
		Success: true,
		Friends: friends,
		Message: "",
	})
}

func (h *FriendHandler) GetFriendRequests(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetFriendRequestsResponse{
			Success:  false,
			Requests: nil,
			Message:  "User ID is required",
		})
	}

	requestType := c.Query("type", "incoming")
	if requestType != "incoming" && requestType != "outgoing" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetFriendRequestsResponse{
			Success:  false,
			Requests: nil,
			Message:  "Invalid request type. Must be 'incoming' or 'outgoing'",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.GetFriendRequestsResponse{
			Success:  false,
			Requests: nil,
			Message:  "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetFriendRequests(ctx, &proto.GetFriendRequestsRequest{
		UserId: userID,
		Type:   requestType,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetFriendRequestsResponse{
			Success:  false,
			Requests: nil,
			Message:  "Failed to get friend requests: " + err.Error(),
		})
	}

	requests := make([]dto.FriendRequest, len(resp.Requests))
	for i, req := range resp.Requests {
		requests[i] = dto.FriendRequest{
			ID:           req.Id,
			SenderID:     req.SenderId,
			SenderHandle: req.SenderHandle,
			ReceiverID:   req.ReceiverId,
			Status:       req.Status,
			CreatedAt:    req.CreatedAt.String(),
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetFriendRequestsResponse{
		Success:  true,
		Requests: requests,
		Message:  "",
	})
}
