package handlers

import (
	"backend/services"
	"context"

	"github.com/gofiber/fiber/v2"
)

func FirstTimeHandler(c *fiber.Ctx) error {

	users, err := services.ClientDB.GetAllUsers(context.Background())
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return c.JSON(fiber.Map{
		"firstTime": len(users) == 1,
	})
}
