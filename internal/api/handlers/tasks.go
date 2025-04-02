package handlers

import (
	"context"
	"fmt"
	"log"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskHandler struct {
	BackendClient proto.BackendRequestsClient
}

func (h *TaskHandler) SubmitTask(c *fiber.Ctx) error {
	var request dto.TaskSubmissionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.TaskSubmissionResponse{
			Success: false,
			TaskID:  "",
			Message: "Invalid request format: " + err.Error(),
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var deadlineProto *timestamppb.Timestamp
	if request.Deadline != nil && *request.Deadline != "" {
		deadline, err := time.Parse(time.RFC3339, *request.Deadline)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.TaskSubmissionResponse{
				Success: false,
				TaskID:  "",
				Message: "Invalid deadline format: " + err.Error(),
			})
		}
		deadlineProto = timestamppb.New(deadline)
	}

	var customHours int32
	if request.CustomHours != nil {
		customHours = int32(*request.CustomHours)
	}

	var group, groupId string
	if request.Group != nil {
		group = *request.Group
	}
	if request.GroupID != nil {
		groupId = *request.GroupID
	}

	var flagColor, flagName string
	if request.FlagColor != nil {
		flagColor = *request.FlagColor
	}
	if request.FlagName != nil {
		flagName = *request.FlagName
	}

	task := &proto.Task{
		Id:             uuid.New().String(),
		UserId:         request.UserID,
		Name:           request.Name,
		Description:    request.Description,
		Deadline:       deadlineProto,
		Author:         request.Author,
		Group:          group,
		GroupId:        groupId,
		AssignedTo:     request.AssignedTo,
		TaskDifficulty: request.Difficulty,
		CustomHours:    customHours,
		Completion: &proto.CompletionStatus{
			IsCompleted: false,
			ProofUrl:    "",
		},
		Privacy: &proto.PrivacySettings{
			Level:     request.PrivacyLevel,
			ExceptIds: request.PrivacyExceptIDs,
		},
		FlagStatus:   request.FlagStatus,
		FlagColor:    flagColor,
		FlagName:     flagName,
		DisplayOrder: int32(request.DisplayOrder),
	}

	resp, err := h.BackendClient.CreateTask(ctx, &proto.CreateTaskRequest{
		Task: task,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.TaskSubmissionResponse{
			Success: false,
			TaskID:  "",
			Message: "Failed to create task: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.TaskSubmissionResponse{
		Success: true,
		TaskID:  resp.TaskId,
		Message: "Task created successfully",
	})
}

func (h *TaskHandler) SubmitTasksBatch(c *fiber.Ctx) error {
	var request dto.BatchTaskSubmissionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.BatchTaskSubmissionResponse{
			Success: false,
			TaskIDs: []string{},
			Message: "Invalid request format: " + err.Error(),
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	tasks := make([]*proto.Task, len(request.Tasks))
	for i, taskSubmission := range request.Tasks {
		var deadlineProto *timestamppb.Timestamp
		if taskSubmission.Deadline != nil && *taskSubmission.Deadline != "" {
			deadline, err := time.Parse(time.RFC3339, *taskSubmission.Deadline)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(dto.BatchTaskSubmissionResponse{
					Success: false,
					TaskIDs: []string{},
					Message: fmt.Sprintf("Invalid deadline format for task %d: %v", i, err),
				})
			}
			deadlineProto = timestamppb.New(deadline)
		}

		var customHours int32
		if taskSubmission.CustomHours != nil {
			customHours = int32(*taskSubmission.CustomHours)
		}

		var flagColor, flagName string
		if taskSubmission.FlagColor != nil {
			flagColor = *taskSubmission.FlagColor
		}
		if taskSubmission.FlagName != nil {
			flagName = *taskSubmission.FlagName
		}

		tasks[i] = &proto.Task{
			Id:             taskSubmission.Id,
			UserId:         taskSubmission.UserID,
			Name:           taskSubmission.Name,
			Description:    taskSubmission.Description,
			Deadline:       deadlineProto,
			Author:         taskSubmission.Author,
			Group:          getStringValue(taskSubmission.Group),
			GroupId:        getStringValue(taskSubmission.GroupID),
			AssignedTo:     taskSubmission.AssignedTo,
			TaskDifficulty: taskSubmission.Difficulty,
			CustomHours:    customHours,
			Completion: &proto.CompletionStatus{
				IsCompleted: false,
				ProofUrl:    "",
			},
			Privacy: &proto.PrivacySettings{
				Level:     taskSubmission.PrivacyLevel,
				ExceptIds: taskSubmission.PrivacyExceptIDs,
			},
			ProofNeeded:  taskSubmission.ProofNeeded,
			FlagStatus:   taskSubmission.FlagStatus,
			FlagColor:    flagColor,
			FlagName:     flagName,
			DisplayOrder: int32(taskSubmission.DisplayOrder),
		}
	}

	resp, err := h.BackendClient.CreateTasksBatch(ctx, &proto.CreateTasksBatchRequest{
		Tasks: tasks,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.BatchTaskSubmissionResponse{
			Success: false,
			TaskIDs: []string{},
			Message: "Failed to create tasks: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.BatchTaskSubmissionResponse{
		Success: true,
		TaskIDs: resp.TaskIds,
		Message: fmt.Sprintf("Successfully created %d tasks", len(resp.TaskIds)),
	})
}

func (h *TaskHandler) GetUserTasks(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
			"tasks":   []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
			"tasks":   []interface{}{},
		})
	}

	requesterID := c.Query("requester_id", "")

	log.Printf("REST API: GetUserTasks for userID=%s, requesterID=%s", userID, requesterID)

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUserTasks(ctx, &proto.GetUserTasksRequest{
		UserId:      userID,
		RequesterId: requesterID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch tasks: " + err.Error(),
			"tasks":   []interface{}{},
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
			"tasks":   []interface{}{},
		})
	}

	tasks := make([]map[string]interface{}, len(resp.Tasks))
	for i, task := range resp.Tasks {
		var deadline *string
		if task.Deadline != nil {
			deadlineStr := task.Deadline.AsTime().Format(time.RFC3339)
			deadline = &deadlineStr
		}

		tasks[i] = map[string]interface{}{
			"id":                 task.Id,
			"user_id":            task.UserId,
			"name":               task.Name,
			"description":        task.Description,
			"created_at":         task.CreatedAt.AsTime().Format(time.RFC3339),
			"deadline":           deadline,
			"author":             task.Author,
			"group":              task.Group,
			"group_id":           task.GroupId,
			"assigned_to":        task.AssignedTo,
			"task_difficulty":    task.TaskDifficulty,
			"custom_hours":       task.CustomHours,
			"is_completed":       task.Completion.IsCompleted,
			"proof_url":          task.Completion.ProofUrl,
			"privacy_level":      task.Privacy.Level,
			"privacy_except_ids": task.Privacy.ExceptIds,
			"flag_status":        task.FlagStatus,
			"flag_color":         task.FlagColor,
			"flag_name":          task.FlagName,
			"display_order":      task.DisplayOrder,
		}
	}

	log.Printf("REST API: Returning %d tasks for userID=%s", len(tasks), userID)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"tasks":   tasks,
		"message": nil,
	})
}

func (h *TaskHandler) UpdateTask(c *fiber.Ctx) error {
	var request dto.TaskUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.TaskUpdateResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	var deadline *timestamppb.Timestamp
	if request.Deadline != nil && *request.Deadline != "" {
		parsedDeadline, err := time.Parse(time.RFC3339, *request.Deadline)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(dto.TaskUpdateResponse{
				Success: false,
				Message: "Invalid deadline format: " + err.Error(),
			})
		}
		deadline = timestamppb.New(parsedDeadline)
	}

	var assignedTo []string
	if request.AssignedTo == nil {
		assignedTo = []string{}
	} else {
		assignedTo = request.AssignedTo
	}

	var privacyExceptIds []string
	if request.PrivacyExceptIDs == nil {
		privacyExceptIds = []string{}
	} else {
		privacyExceptIds = request.PrivacyExceptIDs
	}

	var customHours int32
	if request.CustomHours != nil {
		customHours = int32(*request.CustomHours)
	}

	var flagColor string
	if request.FlagColor != nil {
		flagColor = *request.FlagColor
	}

	var flagName string
	if request.FlagName != nil {
		flagName = *request.FlagName
	}

	var proofDescription string
	if request.ProofDescription != nil {
		proofDescription = *request.ProofDescription
	}

	task := &proto.Task{
		Id:             request.ID,
		UserId:         request.UserID,
		Name:           request.Name,
		Description:    request.Description,
		Deadline:       deadline,
		AssignedTo:     assignedTo,
		TaskDifficulty: request.Difficulty,
		CustomHours:    customHours,
		Completion: &proto.CompletionStatus{
			IsCompleted:       request.IsCompleted,
			ProofUrl:          request.ProofURL,
			NeedsConfirmation: request.RequiresConfirmation,
		},
		Privacy: &proto.PrivacySettings{
			Level:     request.PrivacyLevel,
			ExceptIds: privacyExceptIds,
		},
		ProofNeeded:      request.ProofNeeded,
		ProofDescription: proofDescription,
		FlagStatus:       request.FlagStatus,
		FlagColor:        flagColor,
		FlagName:         flagName,
		DisplayOrder:     int32(request.DisplayOrder),
	}

	resp, err := h.BackendClient.UpdateTask(ctx, &proto.UpdateTaskRequest{
		Task: task,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.TaskUpdateResponse{
			Success: false,
			Message: "Failed to update task: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.TaskUpdateResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.TaskUpdateResponse{
		Success: true,
		Message: "Task updated successfully",
	})
}

func (h *TaskHandler) UpdateTaskOrder(c *fiber.Ctx) error {
	var request dto.TaskOrderUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.TaskOrderUpdateResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	taskOrderItems := make([]*proto.TaskOrderItem, len(request.Tasks))
	for i, item := range request.Tasks {
		taskOrderItems[i] = &proto.TaskOrderItem{
			TaskId:       item.TaskID,
			DisplayOrder: int32(item.DisplayOrder),
		}
	}

	resp, err := h.BackendClient.UpdateTaskOrder(ctx, &proto.UpdateTaskOrderRequest{
		UserId: request.UserID,
		Tasks:  taskOrderItems,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.TaskOrderUpdateResponse{
			Success: false,
			Message: "Failed to update task order: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.TaskOrderUpdateResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.TaskOrderUpdateResponse{
		Success: true,
		Message: "Task order updated successfully",
	})
}

func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func (h *TaskHandler) ConfirmTaskCompletion(c *fiber.Ctx) error {
	var request dto.ConfirmTaskCompletionRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request format: " + err.Error(),
		})
	}

	if request.TaskID == "" || request.ConfirmerID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Task ID and confirmer ID are required",
		})
	}

	token := c.Get("Authorization")
	if token == "" && request.Token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
		})
	}

	if token == "" {
		token = request.Token
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.ConfirmTaskCompletion(ctx, &proto.ConfirmTaskCompletionRequest{
		TaskId:      request.TaskID,
		ConfirmerId: request.ConfirmerID,
		IsConfirmed: request.IsConfirmed,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to confirm task completion: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (h *TaskHandler) GetUsersTasksBatch(c *fiber.Ctx) error {
	var request dto.GetUsersTasksBatchRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersTasksBatchResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if len(request.UserIds) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersTasksBatchResponse{
			Success: false,
			Message: "User IDs are required",
		})
	}

	token := request.Token
	if token == "" {
		token = c.Get("Authorization")
	}

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.GetUsersTasksBatchResponse{
			Success: false,
			Message: "Authorization token is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendClient.GetUsersTasksBatch(ctx, &proto.GetUsersTasksBatchRequest{
		UserIds:     request.UserIds,
		RequesterId: request.RequesterId,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.GetUsersTasksBatchResponse{
			Success: false,
			Message: "Failed to fetch users tasks: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.GetUsersTasksBatchResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	// Convert gRPC response to REST DTO
	userTasks := make(map[string][]dto.TaskResponse)

	for userId, userTaskData := range resp.UserTasks {
		tasks := make([]dto.TaskResponse, len(userTaskData.Tasks))

		for i, task := range userTaskData.Tasks {
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

		userTasks[userId] = tasks
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetUsersTasksBatchResponse{
		Success:   true,
		UserTasks: userTasks,
	})
}
