package main

import (
	"backend/database"
	"backend/model"
	"backend/controller"
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

	db := database.CreateConnection(ctx)
	defer db.Close(ctx)

	appctx := model.AppContext{
		DB: db,
	}

	app := fiber.New()
	app.Use(logger.New())

	index := app.Group("/api/v1")
	controller.AuthRouter(index, appctx)

	log.Fatal(app.Listen(":42069"))
}
