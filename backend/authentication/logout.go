package authentication

import (
	"github.com/gofiber/fiber/v2"
)

func Logout(ctx *fiber.Ctx) error {
	err = RemoveSessionAndCookie(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully logged out!"})
}
