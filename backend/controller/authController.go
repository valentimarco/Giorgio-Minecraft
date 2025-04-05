package controller

import (
	"backend/controller/handler"
	"backend/model"

	"github.com/gofiber/fiber/v2"
)


func AuthRouter(app fiber.Router, appctx model.AppContext){
	app.Post("register", handler.RegisterHandler(appctx))
	app.Get("first-time", handler.FirstTimeHandler(appctx))
}