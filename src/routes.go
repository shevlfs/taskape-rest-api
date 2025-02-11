package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func ping(c fiber.Ctx) error {
	return c.SendString(time.Now().UTC().Format(time.RFC850))
}


func verificationCodeRequestRoute(c fiber.Ctx) error {
    var phone = c.Query("phone");
    


}

func checkVerificationCode(c fiber.Ctx) error {

}
