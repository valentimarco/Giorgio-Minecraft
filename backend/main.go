package main

import (
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
	

	log.Fatal(app.Listen(":42069"))
}
