package communication

import (
	"time"

	"go-chat-app/authentication"
	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type ChatRoom struct {
	ID          int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string        `json:"title" gorm:"not null"`
	CreatedAt   time.Time     `json:"createdat" gorm:"not null;default:CURRENT_TIMESTAMP"`
	Description string        `json:"description"`
	Members     pq.Int64Array `json:"members" gorm:"type:integer[]"`
	Admins      pq.Int64Array `json:"admins" gorm:"not null;type:integer[]"`
	Owner       int           `json:"owner" gorm:"not null"`
	Messages    pq.Int64Array `json:"messages" gorm:"type:integer[]"`
	Type        string        `json:"type"`
}

func GetChatRooms(ctx *fiber.Ctx) error {
	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

  if userID == -1 {
    ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
    return nil
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
  
  if userID == -1 {
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

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

  if userID == -1 {
    ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
    return nil
  }

	body.Owner = userID
	body.Members = append(body.Members, int64(userID))
	body.Admins = append(body.Admins, int64(userID))

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

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

  if userID == -1 {
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

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

  if userID == -1 {
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

func GetUsers(ctx *fiber.Ctx) error {
	var response []authentication.User

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}
  
  if userID == -1 {
    ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
    return nil
  }

	err = database.DB.Table("users").Select("id", "username", "about").Where("NOT id = ?", userID).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetUserByID(ctx *fiber.Ctx) error {
	var response authentication.User

	var id = ctx.Params("id")

	err := database.DB.Table("users").Select("id", "username", "about").Where("id = ?", id).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetUserData(ctx *fiber.Ctx) error {
	var response authentication.User

	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}
  
  if userID == -1 {
    ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
    return nil
  }

	err = database.DB.Table("users").Where("id = ?", userID).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetChatRoomMembers(ctx *fiber.Ctx) error {
	var response []authentication.User
	ID := ctx.Params("id")

	var membersIDs ChatRoom
	err := database.DB.Table("chat_rooms").Select("members").Where("id = ?", ID).First(&membersIDs).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = database.DB.Table("users").Where("id = ANY(?)", membersIDs.Members).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}
