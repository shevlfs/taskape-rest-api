package routes

import (
	"context"
	"fmt"
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
		panic(err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

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

func (h *Handler) getAuthTokens(phone string) (string, string, bool) {
	response, err := h.BackendRequestsClient.LoginNewUser(context.Background(), &pb.NewUserLoginRequest{Phone: phone})
	if err != nil {
		print(err)
		return "", "", false
	}
	return response.Token, response.RefreshToken, response.ProfileExists
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

	authToken, refreshToken, profileExists := h.getAuthTokens(phone)

	println("code accepted & tokens given out to " + phone)

	return c.JSON(fiber.Map{
		"authToken":     authToken,
		"refreshToken":  refreshToken,
		"profileExists": profileExists,
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

	task := &pb.Task{
		Id:             uuid.New().String(),
		UserId:         request.UserID,
		Name:           request.Name,
		Description:    request.Description,
		Deadline:       deadlineProto,
		Author:         request.Author,
		Group:          *request.Group,
		GroupId:        *request.GroupID,
		AssignedTo:     request.AssignedTo,
		TaskDifficulty: request.Difficulty,
		CustomHours:    int32(*request.CustomHours),
		Completion: &pb.CompletionStatus{
			IsCompleted: false,
			ProofUrl:    "",
		},
		Privacy: &pb.PrivacySettings{
			Level:     request.PrivacyLevel,
			ExceptIds: request.PrivacyExceptIDs,
		},
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

		tasks[i] = &pb.Task{
			Id:             uuid.New().String(),
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

	resp, err := h.BackendRequestsClient.GetUserTasks(ctx, &pb.GetUserTasksRequest{
		UserId: userID,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch tasks: " + err.Error(),
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
			ID:               task.Id,
			UserID:           task.UserId,
			Name:             task.Name,
			Description:      task.Description,
			CreatedAt:        task.CreatedAt.AsTime().Format(time.RFC3339),
			Deadline:         deadline,
			Author:           task.Author,
			Group:            task.Group,
			GroupID:          task.GroupId,
			AssignedTo:       task.AssignedTo,
			TaskDifficulty:   task.TaskDifficulty,
			CustomHours:      int(task.CustomHours),
			IsCompleted:      task.Completion.IsCompleted,
			ProofURL:         task.Completion.ProofUrl,
			PrivacyLevel:     task.Privacy.Level,
			PrivacyExceptIDs: task.Privacy.ExceptIds,
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"tasks":   tasks,
	})
}
