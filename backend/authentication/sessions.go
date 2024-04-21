package authentication

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"go-chat-app/database"

	"github.com/gofiber/fiber/v2"
)

var errExpiredCookie = errors.New("expiredCookie")

type HasCookie interface {
	Cookies(string, ...string) string
}

type Session struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userid" gorm:"not null"`
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

func AddSessionToDB(ID string, UserID string) error {
	var oneWeek = 7 * 24 * time.Hour
	return database.DB.Table("sessions").Create(&Session{ID: ID, UserID: UserID, ExpiresAt: time.Now().Add(oneWeek)}).Error
}

func RemoveSessionFromDB(ID string) error {
	var session Session

	err := database.DB.Table("sessions").Where("id = ?", ID).First(&session).Error
	if err != nil {
		return err
	}

	if session.UserID != "" {
		err = database.DB.Table("users").Where("id = ?", session.UserID).Select("email_verified").Updates(&User{EmailVerified: false}).Error
		if err != nil {
			return err
		}
	}

	return database.DB.Table("sessions").Where("id = ?", ID).Delete(&Session{}).Error
}

func GetUserIDFromSession(ctx HasCookie) (string, error) {
	var cookie = ctx.Cookies("session_cookie")
	if cookie == "" {
		return "", http.ErrNoCookie
	}

	var session Session

	err := database.DB.Table("sessions").Where("id = ?", cookie).Select("expires_at", "user_id").First(&session).Error
	if err != nil {
		return "", err
	}

	if session.ExpiresAt.Before(time.Now()) {
		err = database.DB.Table("sessions").Where("id = ?", cookie).Delete(&Session{}).Error
		if err != nil {
			return "", err
		}

		err = database.DB.Table("users").Where("user_id = ?", session.UserID).Select("email_verified").Updates(&User{EmailVerified: false}).Error
		if err != nil {
			return "", err
		}

		return "", errExpiredCookie
	}

	return session.UserID, err
}

func RemoveSessionAndCookie(ctx *fiber.Ctx) error {
	var cookie = ctx.Cookies("session_cookie")
	err := RemoveSessionFromDB(cookie)

	ctx.Cookie(&fiber.Cookie{
		Name:     "session_cookie",
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   1,
		Expires:  time.Now(),
		Secure:   true,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return err
}
