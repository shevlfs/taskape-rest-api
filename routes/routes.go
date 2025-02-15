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
