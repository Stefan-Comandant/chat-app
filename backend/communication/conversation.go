package communication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
) 

type Conversation struct {
	ID		 int		   `json:"id" gorm:"primaryKey;autoIncrement"`
	Messages pq.Int64Array `json:"messages" gorm:"type:integer[]"`
	Members	 pq.Int64Array `json:"members" gorm:"type:integer[]"`
}

func GetConversationByID(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{ "status": "success", "response": "" })
}

func GetConversations(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{ "status": "success", "response": "" })
}

func CreateConversation(ctx *fiber.Ctx) error {
	// Start a conversation with another user

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{ "status": "success", "response": "Succesfully added conversation!" })
}

func DeleteConversation(ctx *fiber.Ctx) error {
	// Delete the conversation and the resective messages

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{ "status": "success", "response": "Succesfully deleted conversation!" })
}

