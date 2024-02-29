package main

import (
	"go-chat-app/authentication"
	"go-chat-app/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	database.DB.Table("users").AutoMigrate(&authentication.User{})
	database.DB.Table("sessions").AutoMigrate(&authentication.Session{})

	if err != nil {
		log.Fatalf("Error while opening db: \n%v\n", err)
	}

	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type",
		AllowCredentials: true,
	}))
	router.Use(func(ctx *fiber.Ctx) error {
		log.Printf("New request To %v with method %v\n", ctx.Path(), ctx.Method())

		return ctx.Next()
	})

	router.Post("/register", authentication.Register)
	router.Post("/login", authentication.Login)
	router.Get("/logout", authentication.Logout)
	router.Get("/code/:code", authentication.EmailCodeVerifier)

	// Start server
	err = router.Listen(":7000")
	if err != nil {
		log.Fatalf("Error while starting server: \n%v\n", err)
	}
}
