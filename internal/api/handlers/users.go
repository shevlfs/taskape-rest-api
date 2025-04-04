package handlers

import (
	"context"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"
	"time"

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



func (h *UserHandler) GetUsersBatch(c *fiber.Ctx) error {
	var request dto.GetUsersBatchRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersBatchResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if len(request.UserIds) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersBatchResponse{
			Success: false,
			Message: "User IDs are required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.GetUsersBatchResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUsersBatch(ctx, &proto.GetUsersBatchRequest{
		UserIds: request.UserIds,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetUsersBatchResponse{
			Success: false,
			Message: "Failed to fetch users: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersBatchResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	
	users := make([]dto.UserResponse, len(resp.Users))
	for i, user := range resp.Users {
		users[i] = dto.UserResponse{
			Id:             user.Id,
			Handle:         user.Handle,
			Bio:            user.Bio,
			ProfilePicture: user.ProfilePicture,
			Color:          user.Color,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetUsersBatchResponse{
		Success: true,
		Users:   users,
	})
}

func (h *UserHandler) EditUserProfile(c *fiber.Ctx) error {
	var request dto.EditUserProfileRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.EditUserProfileResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.UserId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.EditUserProfileResponse{
			Success: false,
			Message: "User ID is required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.EditUserProfileResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.EditUserProfile(ctx, &proto.EditUserProfileRequest{
		UserId:         request.UserId,
		Handle:         request.Handle,
		Bio:            request.Bio,
		Color:          request.Color,
		ProfilePicture: request.ProfilePicture,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.EditUserProfileResponse{
			Success: false,
			Message: "Failed to edit user profile: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.EditUserProfileResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.EditUserProfileResponse{
		Success: true,
	})
}

func (h *UserHandler) CreateGroup(c *fiber.Ctx) error {
	var request dto.CreateGroupRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateGroupResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.CreatorId == "" || request.GroupName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateGroupResponse{
			Success: false,
			Message: "Creator ID and group name are required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateGroupResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.CreateGroup(ctx, &proto.CreateGroupRequest{
		CreatorId:   request.CreatorId,
		GroupName:   request.GroupName,
		Description: request.Description,
		Color:       request.Color,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.CreateGroupResponse{
			Success: false,
			Message: "Failed to create group: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.CreateGroupResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.CreateGroupResponse{
		Success: true,
		GroupId: resp.GroupId,
	})
}

func (h *UserHandler) GetGroupTasks(c *fiber.Ctx) error {
	groupID := c.Params("groupID")
	if groupID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetGroupTasksResponse{
			Success: false,
			Message: "Group ID is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.GetGroupTasksResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	requesterID := c.Query("requester_id", "")

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetGroupTasks(ctx, &proto.GetGroupTasksRequest{
		GroupId:     groupID,
		RequesterId: requesterID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetGroupTasksResponse{
			Success: false,
			Message: "Failed to fetch group tasks: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetGroupTasksResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	
	tasks := make([]dto.TaskResponse, len(resp.Tasks))
	for i, task := range resp.Tasks {
		var deadline *string
		if task.Deadline != nil {
			deadlineStr := task.Deadline.AsTime().Format(time.RFC3339)
			deadline = &deadlineStr
		}

		tasks[i] = dto.TaskResponse{
			ID:                   task.Id,
			UserID:               task.UserId,
			Name:                 task.Name,
			Description:          task.Description,
			CreatedAt:            task.CreatedAt.AsTime().Format(time.RFC3339),
			Deadline:             deadline,
			Author:               task.Author,
			Group:                task.Group,
			GroupID:              task.GroupId,
			AssignedTo:           task.AssignedTo,
			TaskDifficulty:       task.TaskDifficulty,
			CustomHours:          int(task.CustomHours),
			IsCompleted:          task.Completion.IsCompleted,
			ProofURL:             task.Completion.ProofUrl,
			PrivacyLevel:         task.Privacy.Level,
			PrivacyExceptIDs:     task.Privacy.ExceptIds,
			FlagStatus:           task.FlagStatus,
			FlagColor:            &task.FlagColor,
			FlagName:             &task.FlagName,
			DisplayOrder:         int(task.DisplayOrder),
			RequiresConfirmation: task.Completion.NeedsConfirmation,
			IsConfirmed:          task.Completion.IsConfirmed,
		}
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetGroupTasksResponse{
		Success: true,
		Tasks:   tasks,
	})
}

func (h *UserHandler) InviteToGroup(c *fiber.Ctx) error {
	var request dto.InviteToGroupRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.InviteToGroupResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.GroupId == "" || request.InviterId == "" || request.InviteeId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.InviteToGroupResponse{
			Success: false,
			Message: "Group ID, inviter ID, and invitee ID are required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.InviteToGroupResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.InviteToGroup(ctx, &proto.InviteToGroupRequest{
		GroupId:   request.GroupId,
		InviterId: request.InviterId,
		InviteeId: request.InviteeId,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.InviteToGroupResponse{
			Success: false,
			Message: "Failed to send group invitation: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.InviteToGroupResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.InviteToGroupResponse{
		Success:  true,
		InviteId: resp.InviteId,
	})
}

func (h *UserHandler) AcceptGroupInvite(c *fiber.Ctx) error {
	var request dto.AcceptGroupInviteRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.AcceptGroupInviteResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.InviteId == "" || request.UserId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.AcceptGroupInviteResponse{
			Success: false,
			Message: "Invite ID and user ID are required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.AcceptGroupInviteResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.AcceptGroupInvite(ctx, &proto.AcceptGroupInviteRequest{
		InviteId: request.InviteId,
		UserId:   request.UserId,
		Accept:   request.Accept,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.AcceptGroupInviteResponse{
			Success: false,
			Message: "Failed to process group invitation: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.AcceptGroupInviteResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.AcceptGroupInviteResponse{
		Success: true,
	})
}

func (h *UserHandler) KickUserFromGroup(c *fiber.Ctx) error {
	var request dto.KickUserFromGroupRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.KickUserFromGroupResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.GroupId == "" || request.AdminId == "" || request.UserId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.KickUserFromGroupResponse{
			Success: false,
			Message: "Group ID, admin ID, and user ID are required",
		})
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.KickUserFromGroupResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.KickUserFromGroup(ctx, &proto.KickUserFromGroupRequest{
		GroupId: request.GroupId,
		AdminId: request.AdminId,
		UserId:  request.UserId,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.KickUserFromGroupResponse{
			Success: false,
			Message: "Failed to kick user from group: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.KickUserFromGroupResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.KickUserFromGroupResponse{
		Success: true,
	})
}

func (h *UserHandler) GetUserStreak(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUserStreak(ctx, &proto.GetUserStreakRequest{
		UserId: userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch user streak: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
		})
	}

	
	var lastCompletedDate, streakStartDate *string
	if resp.Streak.LastCompletedDate != nil {
		dateStr := resp.Streak.LastCompletedDate.AsTime().Format(time.RFC3339)
		lastCompletedDate = &dateStr
	}

	if resp.Streak.StreakStartDate != nil {
		dateStr := resp.Streak.StreakStartDate.AsTime().Format(time.RFC3339)
		streakStartDate = &dateStr
	}

	return c.Status(fiber.StatusOK).JSON(dto.UserStreakResponse{
		Success:           true,
		CurrentStreak:     resp.Streak.CurrentStreak,
		LongestStreak:     resp.Streak.LongestStreak,
		LastCompletedDate: lastCompletedDate,
		StreakStartDate:   streakStartDate,
	})
}
