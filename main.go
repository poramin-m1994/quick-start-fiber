package main

import (
	"log"
	"quickstart/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// กำหนดเราท์เตอร์
	routers.SetupRoutes(app)

	// เริ่มเซิร์ฟเวอร์
	log.Fatal(app.Listen(":3000"))

}
