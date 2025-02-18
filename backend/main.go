package main

import (
	"backend/handlers"
	"backend/services"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()

	conn := services.CreateConnection(ctx)
	defer conn.Close(ctx)

	app := fiber.New()
	app.Use(logger.New())

	index := app.Group("/api/v1")
	index.Post("register", handlers.RegisterHandler)
	index.Get("first-time", handlers.FirstTimeHandler)
	// controllers.authRouter()

	log.Fatal(app.Listen(":42069"))
}
