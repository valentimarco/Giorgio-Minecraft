package main

import (
	"backend/controller"
	"backend/database"
	_ "backend/docs"
	"backend/services"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


// @title GiorgioMinecraftBackend
// @version 0.0.1
// @description HELL
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:42069
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()

	db := database.CreateConnection(ctx)
	defer db.Close(ctx)

	sm := services.NewServerManager(db)
	defer sm.Close(ctx)

	

	app := fiber.New()
	app.Use(logger.New())

	index := app.Group("/api/v1")
	// index.Get("/docs/*", scalar.HandlerDefault)

	controller.AuthRouter(index, db)
	controller.ServerRouter(index, sm)

	log.Fatal(app.Listen(":42069"))


	// cm.PullImage(context.TODO())
}
