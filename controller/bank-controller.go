package controller

import (
	"api-fiber-mongo/entity"
	"api-fiber-mongo/service"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAccount(c *fiber.Ctx) error {
	accounts := service.GetAccounts()

	return c.Status(fiber.StatusOK).JSON(accounts)
}

func DepositAmount(c *fiber.Ctx) error { //Controller deposit

	valorDeposit := c.Params("balance")

	num, err := strconv.Atoi(valorDeposit)

	result := service.MakeDeposits(num)

	println(result)

	if err != nil {
		fmt.Println("erro ao converter str em int", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "depósito realizado com sucesso",
	})
}

func PostAccount(c *fiber.Ctx) error { //Pega o corpo da requisição, passa pra objeto e executa a função, criando uma conta
	var conta entity.Account
	if err := c.BodyParser(&conta); err != nil {
		log.Println(err.Error())
	}
	service.CreateAccounts(&conta)

	return c.SendStatus(fiber.StatusCreated)
}

func SaqueAmount(c *fiber.Ctx) error { //Controller que realiza o saque
	getValue := c.Params("balance")
	value, err := strconv.Atoi(getValue)

	result := service.MakeWithdraw(value)

	println(result)

	if err != nil {
		fmt.Println("erro ao converter str em int", err)
	}

	notasSomadas := service.Calculo(value)

	showResult := []*entity.Result{
		{Status: "Saque realizado com Sucesso", ValorSaque: value, NotasUsadas: notasSomadas},
	}

	return c.Status(fiber.StatusOK).JSON(showResult)
}
