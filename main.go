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
	err := api.Listen(":3000")
	if err != nil {
		log.Println(err.Error())
	}
}


