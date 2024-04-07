package authentication

import (
  "os"
	"fmt"
  "encoding/base64"

	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
  "gorm.io/gorm/clause"
)

type User struct {
	ID		         int     `json:"id" gorm:"primaryKey;autoIncrement"`
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

  body.Password = hashedPass;

  code, err := GenerateSessionId(6)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }

  var fileType string

  switch string([]byte(body.ProfilePicture)[:15]) {
  case "data:image/png;":
    fileType = "png"
  case "data:image/jpg;":
    fileType = "jpg"
  case "data:image/jpeg":
    fileType = "jpeg"
  }

  fileName := fmt.Sprintf("../profiles/%v.%v", code, fileType)
  

  left, right := 22, 4 

  if fileType == "jpeg" {
    left = 23
  }
  
  fileContent, err := base64.StdEncoding.DecodeString(string([]byte(body.ProfilePicture)[left:len(body.ProfilePicture) - right]))
	if err != nil {
		panic(err)
	}

  err = os.WriteFile(fileName, fileContent, 0755)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }
  
  body.ProfilePicture = code

	err = database.DB.Clauses(clause.Returning{}).Table("users").Create(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

  code, err = GenerateSessionId(6)
  if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }

	// Send email with verification code
	emailBody := fmt.Sprintf("<h1>%v</h1>", code)

	SendGoMail("stefancomandant@gmail.com", body.Email, "", emailBody)
  err = database.DB.Table("verification_sessions").Create(&VerificationSession{Code: code, UserID: body.ID}).Error
   if err != nil {
    ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error() })
    return err
  }
  
  return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully registerd account!", "id": body.ID})
}
