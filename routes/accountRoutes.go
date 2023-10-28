package routes

import (
	"api-fiber-mongo/controller"
	"github.com/gofiber/fiber/v2"
)

func HandleAccountRoutes(api *fiber.App) {
	api.Add(fiber.MethodPost, "/account", controller.CreateAccount)
}
