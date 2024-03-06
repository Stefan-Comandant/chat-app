package authentication

import (
	"fmt"
	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserID   int     `json:"userid" gorm:"primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"not null;"`
	About	 string  `json:"about" gorm:"not null"`
	Email    string  `json:"email" gorm:"not null;unique"`
	Password string  `json:"password" gorm:"not null;"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance" gorm:"not null;default:0"`
}

func Register(ctx *fiber.Ctx) error {
	var body User

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	var mathingEmails int64

	err = database.DB.Table("users").Where("email = ?", body.Email).Count(&mathingEmails).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	// Check if email already exists
	if mathingEmails != 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Email already exists!"})
		return err
	}

	// Send email with verification code

	verificationCode, err := GenerateSessionId(8)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{ "status": "error", "response": err.Error() })
		return err
	}


	emailBody := fmt.Sprintf("<p>Here is your verification code, bitch</p><h1>%v</h1>", verificationCode)

	go CodeTimeOut()
	// Send an email with the verification code
	SendGoMail("stefancomandant@gmail.com", body.Email, "", emailBody)
	emailCodeChannel <- verificationCode

	verificationCodeStatus := <- emailCodeChannel
	if verificationCodeStatus == "failure" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": "Invalid verification code!"})
	}
	if verificationCodeStatus == "timeout" {
		return ctx.Status(fiber.StatusGatewayTimeout).JSON(&fiber.Map{"status": "timeout", "response": "Code verification timeout"})
	}

	//Hash password and store user in db
	hashedPass, err := HashPassword(body.Password)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = database.DB.Table("users").Create(&User{
		Username: body.Username,
		Email:    body.Email,
		Password: hashedPass,
		Currency: body.Currency,
		Balance:  body.Balance,
	}).Error

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully registerd account!" })
}
