package authentication

import (
	"encoding/base64"
	"fmt"
	"os"

	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type User struct {
	ID             string  `json:"id" gorm:"primaryKey"`
	ProfilePicture string  `json:"profilepicture"`
	Username       string  `json:"username" gorm:"not null;"`
	About          string  `json:"about" gorm:"not null"`
	Email          string  `json:"email" gorm:"not null;unique"`
	Password       string  `json:"password" gorm:"not null;"`
	Currency       string  `json:"currency"`
	Balance        float64 `json:"balance" gorm:"not null;default:0"`
	EmailVerified  bool    `json:"emailverified" gorm:"default:f"`
}

func Register(ctx *fiber.Ctx) error {
	var body User

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if len(body.Username) == 0 || len(body.Email) == 0 || len(body.Password) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request body"})
		return nil
	}

	if len(body.ProfilePicture) == 0 {
		body.ProfilePicture = ""
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

	//Hash password and store user in db
	hashedPass, err := HashPassword(body.Password)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	body.Password = hashedPass

	code, err := GenerateSessionId(6)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}
	var fileType string

	body.ID = uuid.NewString()
	if len(body.ProfilePicture) > 0 {
		fileType, err = StoreProfilePicture(body.ProfilePicture, body.ID, "profiles")
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}

		body.ProfilePicture = fileType
	}

	err = database.DB.Clauses(clause.Returning{}).Table("users").Create(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	code, err = GenerateSessionId(6)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	// Send email with verification code
	emailBody := fmt.Sprintf("<h1>%v</h1>", code)

	var session = VerificationSession{
		Code:   code,
		UserID: body.ID,
	}

	SendGoMail("stefancomandant@gmail.com", body.Email, "", emailBody)
	err = database.DB.Clauses(clause.Returning{}).Table("verification_sessions").Create(&session).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	go expireVerificationSession(session)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully registerd account!", "id": body.ID})
}

func StoreProfilePicture(encodedImg string, id string, profileType string) (string, error) {
	var fileType string

	switch string([]byte(encodedImg)[:15]) {
	case "data:image/png;":
		fileType = "png"
	case "data:image/jpg;":
		fileType = "jpg"
	case "data:image/jpeg":
		fileType = "jpeg"
	}

	fileName := fmt.Sprintf("../pictures/%v/%v.%v", profileType, id, fileType)

	left, right := 22, 4

	if fileType == "jpeg" {
		left = 23
	}

	fileContent, err := base64.StdEncoding.DecodeString(string([]byte(encodedImg)[left : len(encodedImg)-right]))
	if err != nil {
		return fileType, err
	}

	return fileType, os.WriteFile(fileName, fileContent, 0755)
}
