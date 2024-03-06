package authentication

import (
	"fmt"
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
	if err != nil || matchingEmails == 0 {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid credentials!"})
		return err
	}

	verificationCode, err := GenerateSessionId(8)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	var emailBody = fmt.Sprintf("<p>Here is your verification code, bitch</p><h1>%v</h1>", verificationCode)

	go CodeTimeOut()
	SendGoMail("stefancomandant@gmail.com", body.Email, "", emailBody)
	emailCodeChannel <- verificationCode

	verificationCodeStatus := <-emailCodeChannel
	if verificationCodeStatus == "failure" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid verification code!"})
	}
	if verificationCodeStatus == "timeout" {
		return ctx.Status(fiber.StatusGatewayTimeout).JSON(&fiber.Map{"status": "timeout", "response": "Code verification timeout"})
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

	Logout(ctx)

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
		SameSite: "lax",
	})
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully logged in account!"})
}
