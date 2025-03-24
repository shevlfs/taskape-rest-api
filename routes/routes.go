package routes

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"taskape-rest-api/dto"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/gofiber/fiber/v2"

	pb "taskape-rest-api/proto"
)

func New() Handler {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	backendHost := os.Getenv("BACKEND_HOST")
	if backendHost == "" {
		backendHost = "localhost:50051"
	}

	conn, err := grpc.NewClient(backendHost, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	conn.Connect()

	client := pb.NewBackendRequestsClient(conn)
	return Handler{BackendRequestsClient: client}
}

type Handler struct {
	pb.BackendRequestsClient
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString(time.Now().UTC().Format(time.RFC850))
}

func (h *Handler) sendCode(phone string) error {
	return nil
}

func (h *Handler) VerificationCodeRequestRoute(c *fiber.Ctx) error {
	var request dto.PhoneVerificationRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	phone := request.Phone
	println(phone)

	if phone == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Phone number is required")
	}
	err := h.sendCode(phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(200)
}

func (h *Handler) checkCode(phone string, code string) bool {
	// gotta move this function & other twilio things to a separate package
	// not doing this right now because this branch is not connected to twilio
	return code == "111111"
}

func (h *Handler) getAuthTokens(phone string) (string, string, bool, int64) {
	response, err := h.BackendRequestsClient.LoginNewUser(context.Background(), &pb.NewUserLoginRequest{Phone: phone})
	if err != nil {
		print(err)
		return "", "", false, -1
	}
	return response.Token, response.RefreshToken, response.ProfileExists, response.UserId
}

func (h *Handler) CheckVerificationCode(c *fiber.Ctx) error {
	var request dto.CheckCodeRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	phone := request.Phone
	code := request.Code

	if phone == "" || code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Phone number and code are required")
	}

	if !h.checkCode(phone, code) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid code")
	}

	authToken, refreshToken, profileExists, userId := h.getAuthTokens(phone)

	println("code accepted & tokens given out to " + phone)

	return c.JSON(fiber.Map{
		"authToken":     authToken,
		"refreshToken":  refreshToken,
		"profileExists": profileExists,
		"userId":        userId,
	})
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
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

	resp, err := h.BackendRequestsClient.GetUser(ctx, &pb.GetUserRequest{
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

	var friends []dto.Friend
	for _, friend := range resp.Friends {
		friends = append(friends, dto.Friend{
			Id:             friend.Id,
			Handle:         friend.Handle,
			ProfilePicture: friend.ProfilePicture,
			Color:          friend.Color,
		})
	}

	var incomingRequests []dto.FriendRequest
	for _, req := range resp.IncomingRequests {
		incomingRequests = append(incomingRequests, dto.FriendRequest{
			Id:           req.Id,
			SenderId:     req.SenderId,
			SenderHandle: req.SenderHandle,
			ReceiverId:   req.ReceiverId,
			Status:       req.Status,
			CreatedAt:    req.CreatedAt.AsTime(),
		})
	}

	var outgoingRequests []dto.FriendRequest
	for _, req := range resp.OutgoingRequests {
		outgoingRequests = append(outgoingRequests, dto.FriendRequest{
			Id:           req.Id,
			SenderId:     req.SenderId,
			SenderHandle: req.SenderHandle,
			ReceiverId:   req.ReceiverId,
			Status:       req.Status,
			CreatedAt:    req.CreatedAt.AsTime(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GetUserResponse{
		Success:          true,
		Id:               resp.Id,
		Handle:           resp.Handle,
		Bio:              resp.Bio,
		ProfilePicture:   resp.ProfilePicture,
		Color:            resp.Color,
		Friends:          friends,
		IncomingRequests: incomingRequests,
		OutgoingRequests: outgoingRequests,
	})
}

func (h *Handler) UpdateTaskOrder(c *fiber.Ctx) error {
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

	taskOrderItems := make([]*pb.TaskOrderItem, len(request.Tasks))
	for i, item := range request.Tasks {
		taskOrderItems[i] = &pb.TaskOrderItem{
			TaskId:       item.TaskID,
			DisplayOrder: int32(item.DisplayOrder),
		}
	}

	resp, err := h.BackendRequestsClient.UpdateTaskOrder(ctx, &pb.UpdateTaskOrderRequest{
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

func (h *Handler) UpdateTask(c *fiber.Ctx) error {
	println("got update task request")
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

	task := &pb.Task{
		Id:             request.ID,
		UserId:         request.UserID,
		Name:           request.Name,
		Description:    request.Description,
		Deadline:       deadline,
		AssignedTo:     assignedTo,
		TaskDifficulty: request.Difficulty,
		CustomHours:    customHours,
		Completion: &pb.CompletionStatus{
			IsCompleted: request.IsCompleted,
			ProofUrl:    request.ProofURL,
		},
		Privacy: &pb.PrivacySettings{
			Level:     request.PrivacyLevel,
			ExceptIds: privacyExceptIds,
		},
		FlagStatus:   request.FlagStatus,
		FlagColor:    flagColor,
		FlagName:     flagName,
		DisplayOrder: int32(request.DisplayOrder),
	}

	resp, err := h.BackendRequestsClient.UpdateTask(ctx, &pb.UpdateTaskRequest{
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

func (h *Handler) VerifyUserToken(c *fiber.Ctx) error {
	var request dto.VerifyTokenRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token := request.Token

	if token == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Token is required")
	}

	resp, err := h.BackendRequestsClient.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if resp.Valid {
		return c.SendStatus(200)
	} else {
		return c.SendStatus(401)
	}
}

func (h *Handler) CheckHandleAvailability(c *fiber.Ctx) error {
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

	resp, err := h.BackendRequestsClient.CheckHandleAvailability(ctx, &pb.CheckHandleRequest{
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

func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	var request dto.RefreshTokenRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token := request.Token
	refreshToken := request.RefreshToken
	phone := request.Phone

	if token == "" || refreshToken == "" || phone == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Token and refresh token are required")
	}

	resp, err := h.BackendRequestsClient.RefreshToken(context.Background(), &pb.RefreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
		Phone:        phone,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"authToken":    resp.Token,
		"refreshToken": resp.RefreshToken,
	})
}

func (h *Handler) RegisterNewProfile(c *fiber.Ctx) error {
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

	response, err := h.BackendRequestsClient.RegisterNewProfile(ctx, &pb.RegisterNewProfileRequest{
		Handle:         request.Handle,
		Bio:            request.Bio,
		Color:          request.Color,
		ProfilePicture: request.ProfilePicture,
		Phone:          request.Phone,
	})

	if err != nil {
		print("error: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"success": response.Success,
		"id":      response.Id,
	})
}

func (h *Handler) SubmitTask(c *fiber.Ctx) error {
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

	task := &pb.Task{
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
		Completion: &pb.CompletionStatus{
			IsCompleted: false,
			ProofUrl:    "",
		},
		Privacy: &pb.PrivacySettings{
			Level:     request.PrivacyLevel,
			ExceptIds: request.PrivacyExceptIDs,
		},
		FlagStatus:   request.FlagStatus,
		FlagColor:    flagColor,
		FlagName:     flagName,
		DisplayOrder: int32(request.DisplayOrder),
	}

	resp, err := h.BackendRequestsClient.CreateTask(ctx, &pb.CreateTaskRequest{
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

func getStringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func (h *Handler) SubmitTasksBatch(c *fiber.Ctx) error {
	var request dto.BatchTaskSubmissionRequest
	println("receiving tasks from client")
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

	tasks := make([]*pb.Task, len(request.Tasks))
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

		tasks[i] = &pb.Task{
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
			Completion: &pb.CompletionStatus{
				IsCompleted: false,
				ProofUrl:    "",
			},
			Privacy: &pb.PrivacySettings{
				Level:     taskSubmission.PrivacyLevel,
				ExceptIds: taskSubmission.PrivacyExceptIDs,
			},
			FlagStatus:   taskSubmission.FlagStatus,
			FlagColor:    flagColor,
			FlagName:     flagName,
			DisplayOrder: int32(taskSubmission.DisplayOrder),
		}
	}

	resp, err := h.BackendRequestsClient.CreateTasksBatch(ctx, &pb.CreateTasksBatchRequest{
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

func (h *Handler) GetUserTasks(c *fiber.Ctx) error {
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

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendRequestsClient.GetUserTasks(ctx, &pb.GetUserTasksRequest{
		UserId: userID,
	})

	if err != nil {
		log.Printf("gRPC error in GetUserTasks: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch tasks: " + err.Error(),
			"tasks":   []interface{}{},
		})
	}

	if !resp.Success {
		log.Printf("Business error in GetUserTasks: %s", resp.Error)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"tasks":   tasks,
		"message": nil,
	})
}

func (h *Handler) SearchUsers(c *fiber.Ctx) error {
	var request dto.SearchUsersRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SearchUsersResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.Query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SearchUsersResponse{
			Success: false,
			Message: "Search query is required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	if request.Limit <= 0 {
		request.Limit = 10
	}

	resp, err := h.BackendRequestsClient.SearchUsers(ctx, &pb.SearchUsersRequest{
		Query: request.Query,
		Limit: request.Limit,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.SearchUsersResponse{
			Success: false,
			Message: "Failed to search users: " + err.Error(),
		})
	}

	var users []dto.UserSearchResult
	for _, user := range resp.Users {
		users = append(users, dto.UserSearchResult{
			Id:             user.Id,
			Handle:         user.Handle,
			ProfilePicture: user.ProfilePicture,
			Color:          user.Color,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.SearchUsersResponse{
		Success: true,
		Users:   users,
	})
}

func (h *Handler) SendFriendRequest(c *fiber.Ctx) error {
	var request dto.SendFriendRequestRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.SenderId == "" || request.ReceiverId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Sender and receiver IDs are required",
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": request.Token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendRequestsClient.SendFriendRequest(ctx, &pb.SendFriendRequestRequest{
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
	})

	if err != nil {
		println("error ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(dto.SendFriendRequestResponse{
			Success: false,
			Message: "Failed to send friend request: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.SendFriendRequestResponse{
		Success:   resp.Success,
		RequestId: resp.RequestId,
		Message:   resp.Error,
	})
}

func (h *Handler) RespondToFriendRequest(c *fiber.Ctx) error {
	var request dto.RespondToFriendRequestRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Invalid request format: " + err.Error(),
		})
	}

	if request.RequestId == "" || request.UserId == "" || request.Response == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Request ID, user ID, and response are required",
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

	resp, err := h.BackendRequestsClient.RespondToFriendRequest(ctx, &pb.RespondToFriendRequestRequest{
		RequestId: request.RequestId,
		UserId:    request.UserId,
		Response:  request.Response,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: "Failed to respond to friend request: " + err.Error(),
		})
	}

	if !resp.Success {
		return c.Status(fiber.StatusBadRequest).JSON(dto.RespondToFriendRequestResponse{
			Success: false,
			Message: resp.Error,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.RespondToFriendRequestResponse{
		Success: true,
		Message: "Friend request " + request.Response + "ed successfully",
	})
}

func (h *Handler) GetUserFriends(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User ID is required",
			"friends": []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Authorization token is required",
			"friends": []interface{}{},
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendRequestsClient.GetUserFriends(ctx, &pb.GetUserFriendsRequest{
		UserId: userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch friends: " + err.Error(),
			"friends": []interface{}{},
		})
	}

	var friends []dto.Friend
	for _, friend := range resp.Friends {
		friends = append(friends, dto.Friend{
			Id:             friend.Id,
			Handle:         friend.Handle,
			ProfilePicture: friend.ProfilePicture,
			Color:          friend.Color,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"friends": friends,
	})
}

func (h *Handler) GetFriendRequests(c *fiber.Ctx) error {
	userID := c.Params("userID")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message":  "User ID is required",
			"requests": []interface{}{},
		})
	}

	requestType := c.Query("type")
	if requestType != "incoming" && requestType != "outgoing" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":  false,
			"message":  "Type must be 'incoming' or 'outgoing'",
			"requests": []interface{}{},
		})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":  false,
			"message":  "Authorization token is required",
			"requests": []interface{}{},
		})
	}

	ctx := context.Background()
	md := metadata.New(map[string]string{
		"authorization": token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := h.BackendRequestsClient.GetFriendRequests(ctx, &pb.GetFriendRequestsRequest{
		UserId: userID,
		Type:   requestType,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success":  false,
			"message":  "Failed to fetch friend requests: " + err.Error(),
			"requests": []interface{}{},
		})
	}

	var requests []dto.FriendRequest
	for _, req := range resp.Requests {
		requests = append(requests, dto.FriendRequest{
			Id:           req.Id,
			SenderId:     req.SenderId,
			SenderHandle: req.SenderHandle,
			ReceiverId:   req.ReceiverId,
			Status:       req.Status,
			CreatedAt:    req.CreatedAt.AsTime(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"requests": requests,
	})
}
