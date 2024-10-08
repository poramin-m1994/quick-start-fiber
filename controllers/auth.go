package controllers

import (
	"quickstart/db"
	"quickstart/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid" // ใช้สำหรับสร้าง UUID
	"golang.org/x/crypto/bcrypt"
)

// GenerateToken สร้าง Token ใหม่สำหรับผู้ใช้
func GenerateToken(userID uint) (string, error) {
	// สร้าง UUID เพื่อเป็น token
	token := uuid.New().String()

	// บันทึก Token ลงในฐานข้อมูล
	newToken := models.Token{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(24 * time.Hour), // ตั้งเวลาหมดอายุเป็น 24 ชั่วโมง
		CreatedAt: time.Now(),
	}

	if err := db.DB.Create(&newToken).Error; err != nil {
		return "", err
	}

	return token, nil
}

// Login ฟังก์ชันสำหรับเข้าสู่ระบบ
func Login(c *fiber.Ctx) error {
	user := new(models.User)
	userToken := ""
	var result = map[string]interface{}{}
	// แปลงข้อมูล JSON ที่รับมาเป็น User struct
	if err := c.BodyParser(user); err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusBadRequest, fiber.StatusBadRequest, "Cannot parse JSON", result)
	}

	// ตรวจสอบผู้ใช้จากฐานข้อมูล
	var foundUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusUnauthorized, fiber.StatusUnauthorized, "User not found", result)
	}

	// ตรวจสอบรหัสผ่าน (เปรียบเทียบรหัสผ่านที่เข้ารหัส)
	if isPass := checkPasswordHash(user.Password, foundUser.Password); !isPass {
		// if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return ResponseJsonWithCOde(c, fiber.StatusUnauthorized, fiber.StatusUnauthorized, "Invalid password", result)
	}

	var token models.Token
	err := db.DB.Where("user_id = ?", foundUser.UserID).Order("created_at desc").First(&token).Error
	// ถ้าไม่มี Token ให้สร้างใหม่
	if err != nil {
		userToken, err = GenerateToken(foundUser.UserID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
		}

	} else {
		userToken = token.Token
	}

	// ถ้า Token มีอยู่แล้ว ให้ตรวจสอบว่า Token หมดอายุหรือไม่
	if token.ExpiresAt.Before(time.Now()) {
		// อัปเดตเวลา ExpiresAt ของ Token ที่มีอยู่ให้เป็น +24 ชั่วโมง
		newExpiresAt := time.Now().Add(24 * time.Hour)
		db.DB.Model(&token).Update("expires_at", newExpiresAt)
		userToken = token.Token
	}

	result = map[string]interface{}{
		"accessToken": userToken,
		// "permission": map[string]bool{
		// 	"openData":     user.OpenDataServiceApproved,
		// 	"notification": user.NotificationServiceApproved,
		// },
	}
	return ResponseJsonWithCOde(c, fiber.StatusOK, 20000, "Success", result)

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
