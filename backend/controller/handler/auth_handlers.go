package handler

import (
	"backend/model/user"
	"backend/util"

	"github.com/gofiber/fiber/v2"
)

// TODO: write a user service to compact all business logic
func RegisterHandler(userService user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		exists, err := userService.ExistUser()
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

		userdto := new(user.UserDTO)
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

		if userdto.Username == "" || userdto.Password == ""{
			c.SendStatus(fiber.StatusUnprocessableEntity)
			return c.JSON(fiber.Map{})
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

		user, err := userService.CreateUser(userdto.Username, passwordHash)
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
