package authentication

import (
	"github.com/gofiber/fiber/v2"
)

func EmailCodeVerifier(ctx *fiber.Ctx) error {
	// Code sent from the user
	var sentCode = ctx.Params("code")

	// Listen for msg with the verification code 
	code := <- emailCodeChannel



	if sentCode == code {
		emailCodeChannel <- "success"
		return nil
	}
	emailCodeChannel <- "failure"
	return nil
}
