package communication

import (
	"log"
	"slices"
	"time"

	"go-chat-app/authentication"
	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ChatRoom struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdat" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Description string         `json:"description"`
	Members     pq.StringArray `json:"members" gorm:"type:text[]"`
	Admins      pq.StringArray `json:"admins" gorm:"not null;type:text[]"`
	Owner       string         `json:"owner" gorm:"not null"`
	Messages    pq.Int64Array  `json:"messages" gorm:"type:integer[]"`
	Type        string         `json:"type"`
}

func GetChatRooms(ctx *fiber.Ctx) error {
	var conversationType = ctx.Params("type")
	if conversationType == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid URL"})
		return nil
	}

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
	}

	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
	}

	if conversationType == "direct" {
		//TODO: add a way to make the websocket conn work
		var IDs []string
		var stuff []string
		err = database.DB.Table("messages").Select(`"from"`, `"to"`).Where(`"from" = ? OR "to" = ?`, userID, userID).Find(&stuff).Error
		if err != nil {
			log.Println(err)
		}

		for _, id := range stuff {
			if !slices.Contains(IDs, id) {
				IDs = append(IDs, id)
			}
		}

		var response []authentication.User

		if len(IDs) == 0 {
			return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
		}

		err = database.DB.Table("users").Where("id IN ?", IDs).Find(&response).Error
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}

		for i, user := range response {
			encoding, err := getProfilePictureEncoding(user)
			if err != nil {
				ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
				return err
			}

			response[i].ProfilePicture = encoding
		}

		return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})

	}

	var response []ChatRoom

	// Select the chat rooms where the userID is present in the members column array of each row
	err = database.DB.Table("chat_rooms").Where("? = ANY(members)", userID).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetChatRoomByID(ctx *fiber.Ctx) error {
	var response ChatRoom
	var id = ctx.Params("id")
	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("chat_rooms").Where("id = ? AND ? = ANY(members)", id, userID).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func CreateChatRoom(ctx *fiber.Ctx) error {
	var body ChatRoom
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if len(body.Title) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request body"})
		return nil
	}

	if len(body.Type) == 0 {
		body.Type = "room"
	}

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	body.Owner = userID
	body.Members = append(body.Members, userID)
	body.Admins = append(body.Admins, userID)
	body.ID = uuid.NewString()

	err = database.DB.Table("chat_rooms").Create(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully created chat room!"})
}

func EditChatRoom(ctx *fiber.Ctx) error {
	var body ChatRoom
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if len(body.ID) == 0 || len(body.Owner) == 0 || len(body.Members) == 0 || len(body.Title) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request body"})
		return nil
	}

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("chat_rooms").Where("id = ? AND ? ANY(members)", body.ID, userID).Save(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully edited chat room!"})
}

func DeleteChatRoom(ctx *fiber.Ctx) error {
	var body ChatRoom
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if len(body.ID) == 0 || len(body.Title) == 0 || len(body.Members) == 0 || len(body.Owner) == 0 {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request body"})
		return nil
	}

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("chat_rooms").Where("id = ?", body.ID).Delete(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully deleted room!"})
}
