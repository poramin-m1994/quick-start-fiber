package routers

import (
	"quickstart/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupPersonalApiRouter(app *fiber.App) {
	// ### v1 group ###
	personalV1 := app.Group("/v1/personal")

	// auth api
	personalAuthV1 := personalV1.Group("/auth")
	personalAuthV1.Post("login", controllers.Login)


	// user api
	personalUserV1 := personalV1.Group("/user")
	//openDataUserV1.Use(filterBearer)
	personalUserV1.Get("/", controllers.GetUsers)
	personalUserV1.Post("/", controllers.CreateUser)

}
