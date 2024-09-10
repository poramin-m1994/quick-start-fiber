package routers

import (
    "quickstart/controllers"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // ตั้งค่าเส้นทางหลัก
    app.Get("/", controllers.Home)
    app.Get("/health-check", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

   SetupOpenDataApiRouter(app)
   SetupPersonalApiRouter(app)
   SetupPortfoliosApiRouter(app)
}
