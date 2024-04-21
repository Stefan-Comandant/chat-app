package middlewares

import (
	"errors"
	"go-chat-app/authentication"
	"go-chat-app/database"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	var path = ctx.Path()
	var method = ctx.Method()

	if strings.Contains(path, "/register") || strings.Contains(path, "/login") || strings.Contains(path, "/code") {
		return ctx.Next()
	}

	userId, err := authentication.GetUserIDFromSession(ctx)
	if err == http.ErrNoCookie {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"status": "error", "response": "Session not present"})
	} else if err != nil {
		ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if strings.Contains(path, "/rooms") && strings.Contains("DELETE, PATCH", method) {
		result := database.DB.Table("chat_rooms").Where("id = ?", ctx.Params("id")).Where("? = ANY(admins) OR ? = owner", userId, userId)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"status": "error", "response": "You don't have the necessary permissions to edit or delete chat room"})
		} else if result.Error != nil {
			ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"status": "error", "response": result.Error.Error()})
			return err
		}
	}

	return ctx.Next()
}
