package communication

import (
	"encoding/json"
	"strconv"
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
	FromID int       `json:"fromid" gorm:"not null"`
	ToID   int       `json:"toid" gorm:"not null"`
	Type   string    `json:"type" gorm:"nont null"`
}

func AddMessage(msg Message) (Message, error) {
	err := database.DB.Table("messages").Create(&msg).Error
	if err != nil {
		return Message{}, err
	}

	var data ChatRoom

	err = database.DB.Clauses(clause.Returning{}).Table("chat_rooms").Where("id = ?", msg.ToID).Find(&data).Error
	if err != nil {
		return Message{}, err
	}

	data.Messages = append(data.Messages, int64(msg.ID))

	err = database.DB.Table("chat_rooms").Where("id = ?", msg.ToID).Save(&data).Error

	return msg, err
}

func GetMessages(ctx *fiber.Ctx) error {
	var response []Message
	var body []int64
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"status": "error", "response": err.Error()})
		return err
	}

	err = database.DB.Table("messages").Where("id IN ?", body).Find(&response).Error
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

		ToID, err := strconv.Atoi(conn.Params("id"))
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		userID, err := authentication.GetUserIDFromSession(conn)
		if err != nil {
			delete(ActiveConnections, conn)
			break
		}

		var message = Message{
			Text:   string(msg.Text),
			FromID: userID,
			ToID:   ToID,
			Type:   msg.Type,
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
