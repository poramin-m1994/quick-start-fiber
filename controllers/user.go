package controllers

import (
	"quickstart/db"
	"quickstart/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser สร้างผู้ใช้ใหม่โดยใช้ Transaction
func CreateUser(c *fiber.Ctx) error {
	var result = map[string]interface{}{}
	// _, shouldReturn, returnValue := GetAuthHeader(c, result)
	// if shouldReturn {
	// 	return returnValue
	// }

	user := new(models.User)

	// แปลงข้อมูล JSON ที่รับมาเป็น User struct
	if err := c.BodyParser(user); err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "Cannot parse JSON", result)
	}
	// เข้ารหัสรหัสผ่าน (Password hashing)
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = hashedPassword

	// เริ่ม Transaction
	tx := db.DB.Begin()

	// ตรวจสอบการเริ่มต้น Transaction
	if tx.Error != nil {
		return ResponseJsonWithCOde(c, fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Failed to start transaction", result)
	}

	// สร้างข้อมูลผู้ใช้ในฐานข้อมูล
	if err := tx.Create(&user).Error; err != nil {
		// Rollback Transaction หากมีข้อผิดพลาด
		tx.Rollback()
		return ResponseJsonWithCOde(c, fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Failed to create user : "+err.Error(), result)
	}

	// Commit Transaction เมื่อไม่มีข้อผิดพลาด
	if err := tx.Commit().Error; err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusInternalServerError, fiber.StatusInternalServerError, "Failed to commit transaction : "+err.Error(), result)
	}

	// ส่งข้อมูลผู้ใช้ที่สร้างสำเร็จกลับไป
	return ResponseJsonWithCOde(c, fiber.StatusOK, 20000, "Success", result)
	// return c.Status(fiber.StatusOK).JSON(user)

}

// GetUsers ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	var result = map[string]interface{}{}

	_, shouldReturn, returnValue := GetAuthHeader(c, result)
	if shouldReturn {
		return returnValue
	}

	db.DB.Find(&users)
	result["users"] = users
	return ResponseJsonWithCOde(c, fiber.StatusOK, 20000, "Success", result)
	// return c.JSON(users)
}

// hashPassword เข้ารหัสรหัสผ่านโดยใช้ bcrypt
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
