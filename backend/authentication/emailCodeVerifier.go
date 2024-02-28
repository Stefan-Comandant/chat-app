package authentication

import (
	"github.com/gofiber/fiber/v2"
)

func EmailCodeVerifier(ctx *fiber.Ctx) error {
	var sentCode = ctx.Params("code")
	code := <- emailCodeChannel

	if sentCode == code {
		emailCodeChannel <- "success"
		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{ "status": "success", "response": "Succesfully registerd account!"})
	}
	emailCodeChannel <- "failure"
	return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{ "status": "error", "response": "Failed to register account!"})
}
