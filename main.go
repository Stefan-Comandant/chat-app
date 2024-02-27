package main

import (
	"go-chat-app/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env vars
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while loading env: \n%v\n", err)
	}

	// Connect to postgresql database
	database.DB, err = database.InitDb(database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode:  os.Getenv("SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	})

	if err != nil {
		log.Fatalf("Error while opening db: \n%v\n", err)
	}

	router := fiber.New()

	// Start server
	err = router.Listen(":7000")
	if err != nil {
		log.Fatalf("Error while starting server: \n%v\n", err)
	}
}
