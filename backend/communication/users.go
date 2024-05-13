package communication

import (
	"encoding/base64"
	"fmt"
	"os"

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

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("users").Select("id", "username", "about", "profile_picture").Where("NOT id = ?", userID).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	for i, user := range response {
		encoding, err := getProfilePictureEncoding(user.ProfilePicture, user.ID, "profiles")
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}
		response[i].ProfilePicture = string(encoding)
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func GetUserByID(ctx *fiber.Ctx) error {
	var response authentication.User

	var id = ctx.Params("id")

	err := database.DB.Table("users").Select("id", "username", "about", "profile_picture").Where("id = ?", id).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	encoding, err := getProfilePictureEncoding(response.ProfilePicture, response.ID, "profiles")
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}
	response.ProfilePicture = string(encoding)

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

	for i, user := range response {
		encoding, err := getProfilePictureEncoding(user.ProfilePicture, user.ID, "profiles")
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
			return err
		}
		response[i].ProfilePicture = string(encoding)
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

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("users").Where("id = ?", userID).First(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	encoding, err := getProfilePictureEncoding(response.ProfilePicture, response.ID, "profiles")
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	response.ProfilePicture = encoding

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func getProfilePictureEncoding(fileType string, id string, profileType string) (string, error) {
	var path = fmt.Sprintf("../pictures/%v/%v.%v", profileType, id, fileType)
	if len(fileType) == 0 || len(id) == 0 {
		fileType = "png"
		path = "../pictures/default.png"
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(content)))
	base64.StdEncoding.Encode(dst, content)

	fullEncoding := fmt.Sprintf("data:image/%v;base64,%v", fileType, string(dst))
	return fullEncoding, err
}
