package controllers

import (
	"quickstart/db"
	"quickstart/models"

	"github.com/gofiber/fiber/v2"
)

// CreateStock สร้างข้อมูลหุ้นใหม่
func CreateStock(c *fiber.Ctx) error {
	var result = map[string]interface{}{}
	stock := new(models.Stock)

	// แปลงข้อมูล JSON ที่รับมาจาก request body
	if err := c.BodyParser(stock); err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "Cannot parse JSON", result)

	}

	// ตรวจสอบว่าสัญลักษณ์หุ้นและชื่อหุ้นถูกต้อง
	if stock.Symbol == "" || stock.Name == "" {
		return ResponseJsonWithCOde(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "Symbol and Name are required", result)
	}

	// บันทึกข้อมูลหุ้นใหม่ในฐานข้อมูล
	if err := db.DB.Create(&stock).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create stock",
		})
	}

	// ส่งข้อมูลหุ้นที่ถูกสร้างกลับไป
	return c.Status(fiber.StatusOK).JSON(stock)
}
