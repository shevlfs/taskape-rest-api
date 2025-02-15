package routes

import (
	"context"
	"time"

	"taskape-rest-api/dto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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

// func (h *Handler) RegisterNewUser(c *fiber.Ctx) error {

// }

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
	return code == "111111"
}

func (h *Handler) getAuthToken(phone string) string {
	response, err := h.BackendRequestsClient.LoginUser(context.Background(), &pb.UserLoginRequest{Phone: phone})
	if err != nil {
		return ""
	}
	return response.Token
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

	authToken := h.getAuthToken(phone)

	println("code accepted for " + phone)

	return c.JSON(fiber.Map{
		"authToken": authToken,
	})
}
