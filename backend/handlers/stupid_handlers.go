package handlers

import (
	"backend/services"
	"context"

	"github.com/gofiber/fiber/v2"
)

func FirstTimeHandler(c *fiber.Ctx) error {

	exists, err := services.ClientDB.ExistUser(context.Background())
	if err != nil {
		c.SendStatus(500)
		return err
	}

	return c.JSON(fiber.Map{
		"firstTime": !exists,
	})
}
