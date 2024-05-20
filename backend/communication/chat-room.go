package communication

import (
	"fmt"
	"slices"
	"time"

	"go-chat-app/authentication"
	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm/clause"
)

type ChatRoom struct {
	ID             string         `json:"id" gorm:"primaryKey"`
	Title          string         `json:"title" gorm:"not null"`
	ProfilePicture string         `json:"profilepicture"`
	CreatedAt      time.Time      `json:"createdat" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Description    string         `json:"description"`
	Members        pq.StringArray `json:"members" gorm:"type:text[]"`
	Admins         pq.StringArray `json:"admins" gorm:"not null;type:text[]"`
	Owner          string         `json:"owner" gorm:"not null"`
	Messages       pq.Int64Array  `json:"messages" gorm:"type:integer[]"`
	Type           string         `json:"type"`
}

func GetChatRooms(ctx *fiber.Ctx) error {
	var conversationType = ctx.Params("type")
	if conversationType == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid URL"})
		return nil
	}

	userID := fmt.Sprintf("%v", ctx.Locals("userID"))

	var response []ChatRoom

	// Select the chat rooms where the userID is present in the members column array of each row
	err := database.DB.Table("chat_rooms").Where("? = ANY(members) AND type = ?", userID, conversationType).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	for i, user := range response {
		encoding, err := getProfilePictureEncoding(user.ProfilePicture, user.ID, "groups")
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}
		response[i].ProfilePicture = string(encoding)
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetChatRoomByID(ctx *fiber.Ctx) error {
	var response ChatRoom
	var id = ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid request"})
	}

	userID := fmt.Sprintf("%v", ctx.Locals("userID"))

	err := database.DB.Table("chat_rooms").Where("id = ? AND ? = ANY(members)", id, userID).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if !slices.Contains(response.Members, userID) {
		return ctx.Status(fiber.StatusNotAcceptable).JSON(&fiber.Map{"status": "error", "response": "You are not part of this group"})
	}

	encoding, err := getProfilePictureEncoding(response.ProfilePicture, response.ID, "groups")
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}
	response.ProfilePicture = string(encoding)

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
		body.Type = "broadcast"
	}

	userID := fmt.Sprintf("%v", ctx.Locals("userID"))

	if body.Type == "broadcast" {
		body.Owner = userID
		body.Admins = append(body.Admins, userID)
	}
	body.Members = append(body.Members, userID)
	body.ID = uuid.NewString()

	if len(body.ProfilePicture) > 0 && body.Type != "direct" {
		fileType, err := authentication.StoreProfilePicture(body.ProfilePicture, body.ID, "groups")
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}

		body.ProfilePicture = fileType
	}

	err = database.DB.Clauses(clause.Returning{}).Table("chat_rooms").Create(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully created chat room!", "id": body.ID})
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

	userID := fmt.Sprintf("%v", ctx.Locals("userID"))

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

	err = database.DB.Table("chat_rooms").Where("id = ?", body.ID).Delete(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully deleted room!"})
}
