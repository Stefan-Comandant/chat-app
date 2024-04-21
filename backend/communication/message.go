package communication

import (
	"encoding/json"
	"time"

	"go-chat-app/authentication"
	"go-chat-app/database"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

var ActiveConnections = make(map[*websocket.Conn]bool)

type Message struct {
	ID     int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Text   string    `json:"text" gorm:"not null"`
	SentAt time.Time `json:"sentat" gorm:"not null;default:CURRENT_TIMESTAMP"`
	From   string    `json:"from" gorm:"not null"`
	To     string    `json:"to" gorm:"not null"`
	Type   string    `json:"type" gorm:"nont null"`
}

func AddMessage(msg Message) (Message, error) {
	err := database.DB.Table("messages").Create(&msg).Error
	if err != nil {
		return Message{}, err
	}

	var data ChatRoom

	err = database.DB.Clauses(clause.Returning{}).Table("chat_rooms").Where("id = ?", msg.To).Find(&data).Error
	if err != nil {
		return Message{}, err
	}

	data.Messages = append(data.Messages, int64(msg.ID))

	err = database.DB.Table("chat_rooms").Where("id = ?", msg.To).Save(&data).Error

	return msg, err
}

func GetMessages(ctx *fiber.Ctx) error {
	userID, err := authentication.GetUserIDFromSession(ctx)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	var response []Message
	var room ChatRoom
	var id = ctx.Params("id")

	err = database.DB.Table("chat_rooms").Where("id = ?", id).First(&room).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = database.DB.Table("messages").Where("id = ANY(?)", room.Messages).Find(&response).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": response})
}

func SendMessage(conn *websocket.Conn) {
	ActiveConnections[conn] = true
	var id = conn.Params("id")

	for {
		dataType, bytesMsg, err := conn.ReadMessage()
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		var msg Message

		err = json.Unmarshal(bytesMsg, &msg)
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		ToID := conn.Params("id")

		userID, err := authentication.GetUserIDFromSession(conn)
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		var message = Message{
			Text: string(msg.Text),
			From: userID,
			To:   ToID,
			Type: msg.Type,
		}

		msg, err = AddMessage(message)
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		jsonData, err := json.Marshal(msg)
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		podcast(jsonData, id, dataType)
	}
}

func DeleteMessage(ctx *fiber.Ctx) error {
	var body Message
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

	if userID == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"status": "error", "response": "Invalid session!"})
		return nil
	}

	err = database.DB.Table("messages").Where("id = ?", body.ID).Delete(&body).Error
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{"status": "success", "response": "Succesfully deleted message!"})
}

func podcast(msg []byte, targetID string, msgType int) {
	for conn := range ActiveConnections {
		var id = conn.Params("id")

		if id != targetID {
			continue
		}
		if err := conn.WriteMessage(msgType, msg); err != nil {
			delete(ActiveConnections, conn)
			continue
		}
	}
}
