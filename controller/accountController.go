package controller

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateAccount(c *fiber.Ctx) error {
	var conta entity.Account
	if err := c.BodyParser(&conta); err != nil {
		log.Println(err.Error())
	}
	service.CreateAccounts(&conta)

	return c.SendStatus(fiber.StatusCreated)
}
