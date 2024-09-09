package controllers

import "github.com/gofiber/fiber/v2"

type ResponseData struct {
	Code           int64       `json:"code"`
	Message        string      `json:"message"`
	ResponseObject interface{} `json:"responseObject"`
}

// Home เป็นคอนโทรลเลอร์ที่ส่งกลับหน้าแรก
func Home(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Welcome to Go Fiber!",
	})
}

func ResponseJsonWithCOde(c *fiber.Ctx, statusCode int, code int64, message string, responseObject interface{}) error {
	if responseObject == nil {
		responseObject = struct{}{}
	}
	result := ResponseData{
		Code:           code,
		Message:        message,
		ResponseObject: responseObject,
	}
	c.Status(int(statusCode))
	return c.JSON(result)
}
