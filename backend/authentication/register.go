package authentication

import (
	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	UserID   int     `json:"userid" gorm:"primaryKey;autoIncrement"`
	Username string  `json:"username" gorm:"not null;"`
	Email    string  `json:"email" gorm:"not null;unique"`
	Password string  `json:"password" gorm:"not null;"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance" gorm:"default:0"`
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

	//TODO: Send email with verification code
	// SendGoMail("")

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

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Sucesfully registred account!"})
}
