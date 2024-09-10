package controllers

import (
	"quickstart/db"
	"quickstart/models"

	"github.com/gofiber/fiber/v2"
)

// CreatePortfolio สร้างพอร์ตโฟลิโอใหม่สำหรับผู้ใช้
func CreatePortfolio(c *fiber.Ctx) error {
	var result = map[string]interface{}{}
	user, shouldReturn, returnValue := GetUserAuthHeader(c, result)
	if shouldReturn {
		return returnValue
	}

	portfolio := new(models.Portfolio)

	// แปลงข้อมูล JSON ที่ส่งเข้ามาเป็น struct Portfolio
	if err := c.BodyParser(portfolio); err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusBadRequest, 40000, "Cannot parse JSON", result)
	}
	portfolio.UserID = user.UserID

	// สร้าง Portfolio ในฐานข้อมูล
	if err := db.DB.Create(&portfolio).Error; err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusInternalServerError, 50000, "Failed to create portfolio", result)
	}

	// ส่งข้อมูล Portfolio ที่ถูกสร้างกลับไป
	return ResponseJsonWithCOde(c, fiber.StatusOK, 20000, "Success", result)

}
