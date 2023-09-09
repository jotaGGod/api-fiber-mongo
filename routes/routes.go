package routes

import (
	"api-fiber-mongo/controller"
	websocket_hub "api-fiber-mongo/websocket-hub"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func HandleRoutes(api *fiber.App) {
	api.Add(fiber.MethodGet, "/account", controller.GetAccount)
	api.Add(fiber.MethodPost, "/account/:balance", controller.DepositAmount)
	api.Add(fiber.MethodDelete, "/account/:balance", controller.SaqueAmount)
	api.Add(fiber.MethodPost, "/account", controller.PostAccount)

	api.Get("/code/:userid/ws", websocket.New(websocket_hub.GenerateCode))
}
