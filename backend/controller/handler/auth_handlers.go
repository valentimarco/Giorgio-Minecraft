package handler

import (
	g "backend/database/gen"
	"backend/model"

	"backend/model/dto"
	"backend/util"

	"github.com/gofiber/fiber/v2"
)
func RegisterHandler(appctx model.AppContext) fiber.Handler{
	return func(c *fiber.Ctx) error {

		exists, err := appctx.DB.ExistUser(c.UserContext())
		if err != nil {
			c.SendStatus(500)
			return c.JSON(fiber.Map{})
		}

		if exists {
			c.SendStatus(209)
			return c.JSON(fiber.Map{
				"type":   "about:blank",
				"title":  "User already exists",
				"status": 209,
				"detail": "There is only and only one user, the admin!",
			}, "problem+json")
		}

		userdto := new(dto.UserDTO)
		if err := c.BodyParser(userdto); err != nil {
			c.SendStatus(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"type":  "https://example.net/validation-error",
				"title": "Your request parameters didn't validate.",
				"invalid-params": []string{
					"",
				},
			})
		}

		passwordHash, err := util.CreateHash(userdto.Password)
		if err != nil {
			return c.JSON(fiber.Map{
				"type":  "",
				"title": "Password Too Long",
				"invalid-params": []string{
					"",
				},
			})
		}

		user, err := appctx.DB.CreateUser(c.UserContext(), g.CreateUserParams{
			Username: userdto.Username,
			Password: passwordHash,
		})
		if err != nil {
			c.SendStatus(fiber.StatusConflict)
			return c.JSON(fiber.Map{})
		}

		token, err := util.GenerateToken(user.ID)
		if err != nil {
			c.SendStatus(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"err": err.Error(),
			})
		}

		c.SendStatus(fiber.StatusOK)

		return c.JSON(fiber.Map{
			"token": token,
		})
	}
}
