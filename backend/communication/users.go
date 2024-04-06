package communication

import (


  "go-chat-app/authentication"
  "go-chat-app/database"

  "github.com/gofiber/fiber/v2"
)

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
