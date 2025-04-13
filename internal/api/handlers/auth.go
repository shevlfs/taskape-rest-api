package handlers

import (
	"context"
	"taskape-rest-api/internal/config"
	"taskape-rest-api/internal/dto"
	proto "taskape-rest-api/proto"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/twilio/twilio-go"
)

type AuthHandler struct {
	BackendClient proto.BackendRequestsClient
	Config        *config.Config
	TwilioClient  *twilio.RestClient
}

func NewAuthHandler(client proto.BackendRequestsClient, cfg *config.Config) *AuthHandler {
	var twilioClient *twilio.RestClient
	if cfg.TwilioAccountSID != "" && cfg.TwilioAuthToken != "" {
		twilioClient = twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: cfg.TwilioAccountSID,
			Password: cfg.TwilioAuthToken,
		})
	}

	return &AuthHandler{
		BackendClient: client,
		Config:        cfg,
		TwilioClient:  twilioClient,
	}
}

func (h *AuthHandler) Ping(c *fiber.Ctx) error {
	return c.SendString(time.Now().UTC().Format(time.RFC850))
}

func (h *AuthHandler) SendVerificationCode(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

const (
	DEV_CODE = "111111"
)

const (
	DEV_ENV = "development"
)

func (h *AuthHandler) CheckVerificationCode(c *fiber.Ctx) error {
	var request dto.CheckCodeRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	phone := request.Phone
	code := request.Code

	if phone == "" || code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Phone number and code are required")
	}
	if h.Config.Environment == DEV_ENV || code == DEV_CODE {

		response, err := h.BackendClient.LoginNewUser(context.Background(), &proto.NewUserLoginRequest{
			Phone: phone,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(dto.VerificationResponse{
			AuthToken:     response.Token,
			RefreshToken:  response.RefreshToken,
			ProfileExists: response.ProfileExists,
			UserId:        response.UserId,
		})
	}

	return c.Status(fiber.StatusBadRequest).SendString("Invalid verification code")
}

func (h *AuthHandler) ValidateToken(c *fiber.Ctx) error {
	var request dto.VerifyTokenRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token := request.Token
	if token == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Token is required")
	}

	resp, err := h.BackendClient.ValidateToken(context.Background(), &proto.ValidateTokenRequest{
		Token: token,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if resp.Valid {
		return c.SendStatus(fiber.StatusOK)
	} else {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var request dto.RefreshTokenRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	token := request.Token
	refreshToken := request.RefreshToken
	phone := request.Phone

	if token == "" || refreshToken == "" || phone == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Token, refresh token, and phone are required")
	}

	resp, err := h.BackendClient.RefreshToken(context.Background(), &proto.RefreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
		Phone:        phone,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(dto.TokenRefreshResponse{
		AuthToken:    resp.Token,
		RefreshToken: resp.RefreshToken,
	})
}
