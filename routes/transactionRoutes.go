package routes

import (
	"api-fiber-mongo/controller"

	"github.com/gofiber/fiber/v2"
)

func HandleTransactionRoutes(api *fiber.App) {
	api.Add(fiber.MethodGet, "/account", controller.GetAccount)
	api.Add(fiber.MethodPost, "/account/:balance", controller.DepositAmount)
	api.Add(fiber.MethodDelete, "/account/:balance", controller.SaqueAmount)
}
