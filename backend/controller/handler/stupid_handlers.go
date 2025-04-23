package handler

import (
	"backend/model/user"

	"github.com/gofiber/fiber/v2"
)

func FirstTimeHandler(userService user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		exists, err := userService.ExistUser()
		if err != nil {
			c.SendStatus(500)
			return err
		}

		return c.JSON(fiber.Map{
			"firstTime": !exists,
		})
	}
}
