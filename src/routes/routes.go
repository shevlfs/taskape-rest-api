package routes

import (
	"time"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func New() Handler {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return Handler{}
}

type Handler struct {
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString(time.Now().UTC().Format(time.RFC850))
}

func (h *Handler) sendCode(phone string) error {
	return nil
}

type Request struct {
	Phone string `json:"phone"`
}

func (h *Handler) VerificationCodeRequestRoute(c *fiber.Ctx) error {
	var request Request
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

func generateAuthToken(phone string) string {
	return "mock-auth-token-for-" + phone
}

type CheckRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func (h *Handler) CheckVerificationCode(c *fiber.Ctx) error {
	var request CheckRequest

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

	authToken := generateAuthToken(phone)

	return c.JSON(fiber.Map{
		"authToken": authToken,
	})
}
