package handler

import (
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

// @Summary Create Server
// @Description Create minecraft server
// @Accept json
// @Produce json
// @Success 200 dto.CreationServer
// @Router /api/v1/create-server
func CreateServer(sm *services.ServerManager) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
		// config := new(servertypes.ServerConfigDTO)
		// if err := c.BodyParser(config); err != nil {
		// 	panic(err)
		// }

		// sm.CreateServer()

		// server, err := appctx.DB.CreateServer(c.UserContext(), gen.CreateServerParams{
		// 	Type:      string(config.Type),
		// 	Name:      config.Name,
		// 	Memory:    config.Memory,
		// 	MaxPlayer: config.MaxPlayers,
		// 	AikaFlag:  config.UseAikaFlag,
		// })
		// if err != nil {
		// 	panic(err)
		// }

		// //set docker label to the container and create the container

		// //TODO: manage whitlist somehow

		// return c.JSON(servertypes.CreationServer{
		// 	ID: server.ID,
		// })
	}
}
