package handlers

import (
	database "backend/database/gen"
	"backend/models/dto"
	"backend/services"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(c *fiber.Ctx) error {

	exists, err := services.ClientDB.ExistUser(c.UserContext())
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

	passwordHash, err := utils.CreateHash(userdto.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"type":  "",
			"title": "Password Too Long",
			"invalid-params": []string{
				"",
			},
		})
	}

	user, err := services.ClientDB.CreateUser(c.UserContext(), database.CreateUserParams{
		Username: userdto.Username,
		Password: passwordHash,
	})
	if err != nil {
		c.SendStatus(fiber.StatusConflict)
		return c.JSON(fiber.Map{})
	}

	token, err := utils.GenerateToken(user.ID)
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
