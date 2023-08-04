package main

import (
	"api-fiber-mongo/database"
	"api-fiber-mongo/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConfigureMongo()

	api := fiber.New()

	routes.HandleRoutes(api)
	if err := api.Listen(":3000"); err != nil {
		log.Println(err.Error())
	}
}
