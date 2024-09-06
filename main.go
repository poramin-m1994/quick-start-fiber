package main

import (
	"log"
	"quickstart/routers"
	"quickstart/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// เชื่อมต่อกับฐานข้อมูล
	db.Connect()

	app := fiber.New()

	// กำหนดเราท์เตอร์
	routers.SetupRoutes(app)

	// เริ่มเซิร์ฟเวอร์
	log.Fatal(app.Listen(":3000"))

}
