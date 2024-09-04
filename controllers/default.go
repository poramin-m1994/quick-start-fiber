package controllers

import "github.com/gofiber/fiber/v2"

// Home เป็นคอนโทรลเลอร์ที่ส่งกลับหน้าแรก
func Home(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "Title": "Welcome to Go Fiber!",
    })
}
