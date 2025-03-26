package handlers

import (
	"context"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
)

type UserHandler struct {
	BackendClient proto.BackendRequestsClient
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "User ID is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"error":   "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUser(ctx, &proto.GetUserRequest{
		UserId: userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch user: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserResponse{
		Success:        true,
		Id:             resp.Id,
		Handle:         resp.Handle,
		Bio:            resp.Bio,
		ProfilePicture: resp.ProfilePicture,
		Color:          resp.Color,
	})
}

func (h *UserHandler) CheckHandleAvailability(c *fiber.Ctx) error {
	var request dto.CheckHandleRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CheckHandleResponse{
			Available: false,
			Message:   "Invalid request format: " + err.Error(),
		})
	}

	if request.Handle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CheckHandleResponse{
			Available: false,
			Message:   "Handle is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	if request.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CheckHandleResponse{
			Available: false,
			Message:   "Token is required",
		})
	}

	handle := request.Handle
	if handle[0] == '@' {
		handle = handle[1:]
	}

	resp, err := h.BackendClient.CheckHandleAvailability(ctx, &proto.CheckHandleRequest{
		Handle: handle,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CheckHandleResponse{
			Available: false,
			Message:   "Failed to check handle availability: " + err.Error(),
		})
	}

	if !resp.Available {
		return c.Status(fiber.StatusOK).JSON(dto.CheckHandleResponse{
			Available: false,
			Message:   "Handle is already taken",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.CheckHandleResponse{
		Available: true,
		Message:   "Handle is available",
	})
}

func (h *UserHandler) RegisterNewProfile(c *fiber.Ctx) error {
	var request dto.RegisterNewProfileRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if request.Phone == "" || request.Token == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Phone and token are required")
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	response, err := h.BackendClient.RegisterNewProfile(ctx, &proto.RegisterNewProfileRequest{
		Handle:         request.Handle,
		Bio:            request.Bio,
		Color:          request.Color,
		ProfilePicture: request.ProfilePicture,
		Phone:          request.Phone,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": response.Success,
		"id":      response.Id,
	})
}
