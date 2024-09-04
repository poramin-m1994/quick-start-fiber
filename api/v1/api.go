package v1

import (
	"github.com/gofiber/fiber/v2"
	// "quickstart/logs"
	// "quickstart/models"
	// "quickstart/services"
)

const (
	// header
	BearerPrefix = "Bearer "

	// response message
	BadRequest                   = "Bad Request"
	BadRequestTh                 = "ข้อมูลไม่ถูกต้อง"
	Unauthorized                 = "Unauthorized"
	SomethingWentWrong           = "Something went wrong"
	SomethingWentWrongTh         = "เกิดข้อผิดพลาด"
	Success                      = "success"
	NotFound                     = "not found"
	UserNotFoundTh               = "ไม่พบข้อมูล"
	SendMailResetPasswordSuccess = "Reset password email has been sent"
	InvalidEmail                 = "Invalid email address"
	InvalidPassword              = "รหัสผ่านต้องมีความยาวอย่างน้อย 8 ตัวอักษร ประกอบด้วยตัวอักษรภาษาอังกฤษอย่างน้อย 1 ตัวและตัวเลขอย่างน้อย 1 ตัว"
	EmailNotVerified             = "Email not verified"
	UserNotApproved              = "User not approved"
	UserHasBeenRejected          = "User has been rejected"
	WrongOldPassword             = "Old password not match"
	DuplicateRecipient           = "Duplicate recipient"
	// date/time format
	DefaultDateTimeFormatResp = "2006/01/02 15:04:05"
	DefaultDateTimeFormatReq  = "2006-01-02 15:04:05"
)

type ResponseData struct {
	Code           int64       `json:"code"`
	Message        string      `json:"message"`
	ResponseObject interface{} `json:"responseObject"`
}

type API struct {
	*fiber.Ctx
}

func (api *API) responseJson(code int64, message string, responseObject interface{}) error {
	result := ResponseData{
		Code:           code,
		Message:        message,
		ResponseObject: responseObject,
	}
	return api.JSON(result)
}

func ResponseJson(c *fiber.Ctx, code int64, message string, responseObject interface{}) error {
	if responseObject == nil {
		responseObject = struct{}{}
	}
	result := ResponseData{
		Code:           code,
		Message:        message,
		ResponseObject: responseObject,
	}
	c.Status(int(code))
	return c.JSON(result)
}

func ResponseJsonWithCode(c *fiber.Ctx, statusCode int, code int64, message string, responseObject interface{}) error {
	if responseObject == nil {
		responseObject = struct{}{}
	}
	result := ResponseData{
		Code:           code,
		Message:        message,
		ResponseObject: responseObject,
	}
	c.Status(statusCode)
	return c.JSON(result)
}

func GetBearer(c *fiber.Ctx) string {
	authorization := getAuthorizationHeader(c)
	n := len(BearerPrefix)
	if len(authorization) < n || authorization[:n] != BearerPrefix {
		return ""
	}
	return authorization[n:]
}

func getAuthorizationHeader(c *fiber.Ctx) string {
	return getHeader(c, fiber.HeaderAuthorization)
}

func getHeader(c *fiber.Ctx, header string) string {
	return c.Get(header)
}

// func GetUserAuthentication(c *fiber.Ctx) (models.User, error) {
// 	token := c.Locals(BearerPrefix).(string)
// 	if token == "" {
// 		token = GetBearer(c)
// 	}
// 	if token == "" {
// 		return models.User{}, errors.New("authorization not found")
// 	}

// 	user, err := services.GetUser(token)
// 	if err != nil {
// 		logs.WithError(err).Error("cannot get user from token: ", token)
// 		return models.User{}, err
// 	}
// 	return user, nil

// }

// func validatePermission(c *fiber.Ctx, permission, action string) bool {
// 	user, err := GetUserAuthentication(c)
// 	if err != nil {
// 		logs.WithError(err).Error("cannot get user")
// 		return false
// 	}
// 	c.Locals("user", user)
// 	if permission == models.PermissionUnrestricted {
// 		return true
// 	}
// 	if !services.ValidatePermission(permission, action, user) {
// 		return false
// 	}
// 	return true
// }

// func HealthCheck(ctx *fiber.Ctx) error {
// 	sql, err := models.DB.DB()
// 	if err != nil {
// 		logs.WithError(err).Error("cannot initial to db")
// 		ctx.Status(500)
// 		return ctx.JSON(fiber.Map{"status": false, "message": "error"})
// 	}
// 	err = sql.Ping()
// 	if err != nil {
// 		logs.WithError(err).Error("cannot connect to db")
// 		ctx.Status(500)
// 		return ctx.JSON(fiber.Map{"status": false, "message": "error"})
// 	}

// 	return ctx.JSON(fiber.Map{"status": true, "message": "success"})
// }
