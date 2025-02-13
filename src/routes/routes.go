package routes

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"

	"github.com/gofiber/fiber/v2"
)

func New() Handler {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return Handler{
		twilioClient: twilio.NewRestClientWithParams(
			twilio.ClientParams{
				Username: os.Getenv("TWILIO_ACCOUNT_SID"),
				Password: os.Getenv("TWILIO_AUTH_TOKEN"),
			},
		),
	}
}

type Handler struct {
	twilioClient *twilio.RestClient
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString(time.Now().UTC().Format(time.RFC850))
}

func (h *Handler) sendCode(phone string) error {
	params := &verify.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	_, err := h.twilioClient.VerifyV2.CreateVerification("VA8cd7f3e1bad0573034eeb9585254d477", params)
	if err != nil {
		return err
	}

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
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)

	resp, err := h.twilioClient.VerifyV2.CreateVerificationCheck("VA8cd7f3e1bad0573034eeb9585254d477", params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Status != nil {
			if *resp.Status == "approved" {
				return true
			}
		}
	}

	return false
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

	cookie := new(fiber.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = authToken
	cookie.HTTPOnly = true
	cookie.Secure = true

	c.Cookie(cookie)

	return c.SendStatus(200)
}
