package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"go-chat-app/database"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userid" gorm:"not null"`
	CreatedAt time.Time `json:"createdat" gorm:"not null;default:CURRENT_TIMESTAMP"`
	ExpiresAt time.Time `json:"expiresat" gorm:"not null"`
}

func GenerateSessionId(length int) (string, error) {
	// Make an empty slice of bytes
	bytes := make([]byte, length)

	// Fill it with random bytes that represent chars
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Turn the slice into a string
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func AddSessionToDB(ID string, UserID int) error {
	var oneWeek = 7 * 24 * time.Hour
	return database.DB.Table("sessions").Create(&Session{ID: ID, UserID: UserID, ExpiresAt: time.Now().Add(oneWeek)}).Error
}

func RemoveSessionFromDB(ID string) error {
	return database.DB.Table("sessions").Where("id = ?", ID).Delete(&Session{}).Error
}

func GetUserIDFromSession(ctx *fiber.Ctx) (int, error) {
	var cookie = ctx.Cookies("session_cookie")
	if cookie == "" {
		return -1, http.ErrNoCookie
	}

	var expiresAt time.Time
	var UserID int

	err := database.DB.Table("sessions").Where("id = ?", cookie).Select("expires_at", "user_id").First(&expiresAt, &UserID).Error

	if expiresAt.Before(time.Now()) {
		return -1, errors.New("expiredCookie")
	}

	return UserID, err
}
func RemoveSessionAndCookie(ctx *fiber.Ctx) error {
	var cookie = ctx.Cookies("session_cookie")
	err := RemoveSessionFromDB(cookie)

	ctx.Cookie(&fiber.Cookie{
		Name:     "session_cookie",
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   -1,
		Expires:  time.Now(),
		Secure:   true,
		HTTPOnly: true,
	})
	return err
}
