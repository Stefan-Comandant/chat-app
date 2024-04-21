package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(ctx *fiber.Ctx) error {
	log.Printf("New request To %v with method %v\n", ctx.Path(), ctx.Method())
	return ctx.Next()
}
