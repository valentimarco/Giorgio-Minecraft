package controller

import (
	"backend/controller/handler"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func ServerRouter(app fiber.Router, sm *services.ServerManager) {


	app.Post("create-server", handler.CreateServer(sm))

	// serverGroup := app.Group(":id")

	// serverGroup.Delete("")
	// serverGroup.Get("")

}
