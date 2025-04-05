package handler

import (
	"backend/model"
	"context"

	"github.com/gofiber/fiber/v2"
)
func FirstTimeHandler(appctx model.AppContext) fiber.Handler{
	return func (c *fiber.Ctx) error {

		exists, err := appctx.DB.ExistUser(context.Background())
		if err != nil {
			c.SendStatus(500)
			return err
		}
	
		return c.JSON(fiber.Map{
			"firstTime": !exists,
		})
	}
}

