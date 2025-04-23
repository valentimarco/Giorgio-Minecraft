package controller

import (
	"backend/controller/handler"
	"backend/database"
	"backend/model/user"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, db *database.DB) {
	userService := user.NewService(db)
	app.Post("register", handler.RegisterHandler(userService))
	app.Get("first-time", handler.FirstTimeHandler(userService))
}
