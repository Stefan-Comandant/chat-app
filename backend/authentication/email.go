package authentication

import (
	"go-chat-app/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
)

func SendGoMail(from string, to string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, from, os.Getenv("APP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

type VerificationSession struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Code   string `json:"code" gorm:"NOT NULL"`
	UserID string `json:"userid" gorm:"NOT NULL"`
}

func VerifyEmailCode(ctx *fiber.Ctx) error {
	var body VerificationSession

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if len(body.Code) == 0 || len(body.UserID) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request body"})
		return nil
	}

	var matchingUsers int64

	err = database.DB.Table("users").Where("id = ?", body.UserID).Count(&matchingUsers).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if matchingUsers == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Email not valid"})
		return nil
	}

	var matchingSessions int64 = 0

	err = database.DB.Table("verification_sessions").Where("user_id = ? AND code = ?", body.UserID, body.Code).Count(&matchingSessions).Error
	if err != nil {
		ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if matchingSessions == 0 {
		ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{"status": "error", "response": "Verification session expired"})
		return nil
	}

	err = database.DB.Table("verification_sessions").Where("user_id = ? AND code = ?", body.UserID, body.Code).Delete(&VerificationSession{}).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = database.DB.Table("users").Where("id = ?", body.UserID).Select("email_verified").Updates(&User{EmailVerified: true}).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = createSession(ctx, body.UserID)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Successfully verified email address!"})
}
