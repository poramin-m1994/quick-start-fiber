package routers

import (
	"quickstart/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPortfoliosApiRouter(app *fiber.App) {
	// ### v1 group ###
	portfolioV1 := app.Group("/v1/portfolios")

	// portfolios api
	portfolioV1.Post("create", controllers.CreatePortfolio)

}
