package routers

import (
	"quickstart/api/v1/opendata"

	"github.com/gofiber/fiber/v2"
)

func SetupOpenDataApiRouter(app *fiber.App) {
	// ### v1 group ###
	openDataV1 := app.Group("/v1/oda")
	// auth api
	openDataAuthV1 := openDataV1.Group("/auth")
	openDataAuthV1.Post("login", opendata.Login)
	openDataAuthV1.Post("logout", opendata.Logout)

	// user api
	openDataUserV1 := openDataV1.Group("/user")
	//openDataUserV1.Use(filterBearer)
	openDataUserV1.Post("/register", opendata.RegisterUser)
	openDataUserV1.Get("/permission", opendata.Permission)
	openDataUserV1.Post("/forgotpassword", opendata.ForgotPassword)
	openDataUserV1.Post("/resetpassword", opendata.NewPassword)
	openDataUserV1.Post("/changepassword", opendata.ChangePassword)
	openDataUserV1.Get("/apikey", opendata.GetAPIKey)
	openDataUserV1.Get("/profile", opendata.Profile)
	openDataUserV1.Delete("/", opendata.DeleteUser)

}
