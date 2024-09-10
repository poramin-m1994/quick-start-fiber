package routers

import (
	"quickstart/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupStocksApiRouter(app *fiber.App) {
	// ### v1 group ###
	stocksV1 := app.Group("/v1/stocks")

	// stocks api
	stocksV1.Post("/create", controllers.CreateStock)

}
