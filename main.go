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

// testes automatiza:
// teste unitario testar funções (Stack)

//----------------------------------
// testes
/*
func TestViewAccount(t *testing.T) {
	app := fiber.New()
	app.Get("/account", viewAccount)

	req := httptest.NewRequest(http.MethodGet, "/account", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
}
*/
