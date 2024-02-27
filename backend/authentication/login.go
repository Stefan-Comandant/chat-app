package authentication

import (
	"go-chat-app/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	var body User
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	var matchingEmails int64

	err = database.DB.Table("users").Where("email = ?", body.Email).Count(&matchingEmails).Error
	if err != nil || matchingEmails != 1 {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid credentials!"})
		return err
	}

	var user User

	err = database.DB.Table("users").Where("email = ?", body.Email).First(&user).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid credentials!"})
		return err
	}

	if !ComparePassword(body.Password, user.Password) {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid credentials!"})
		return err
	}

	sessionID, err := GenerateSessionId(32)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Failed to generate session ID!"})
		return err
	}

	RemoveSessionAndCookie(ctx)

	err = AddSessionToDB(sessionID, user.UserID)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Failed to add session!"})
		return err
	}

	const oneWeek = 60 * 60 * 24 * 7

	ctx.Cookie(&fiber.Cookie{
		Name:     "session_cookie",
		Value:    sessionID,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   oneWeek,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Secure:   true,
		HTTPOnly: true,
	})

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Sucesfully logged in!"})
}
