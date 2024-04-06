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

	code, err := GenerateSessionId(8)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	var emailBody = fmt.Sprintf("<p>Here is your verification code, bitch</p><h1>%v</h1>", code)

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

  err = createSession(ctx, user.ID)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }
	
	SendGoMail("stefancomandant@gmail.com", body.Email, "", emailBody)
  err = database.DB.Table("verification_sessions").Create(&VerificationSession{Code: code, UserID: user.ID}).Error
   if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }


  return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully logged in account!", "id": user.ID})
}

func createSession(ctx *fiber.Ctx, userID int) error {
  sessionID, err := GenerateSessionId(32)
	if err != nil {
		return err
	}

	Logout(ctx)

	err = AddSessionToDB(sessionID, userID)
	if err != nil {
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

  return err
}
