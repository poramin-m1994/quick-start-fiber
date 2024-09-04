package opendata

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	v1 "quickstart/api/v1"
	// "quickstart/internal/data"
	// "quickstart/internal/mail"
	// "quickstart/logs"
	// "quickstart/models"
	// "quickstart/models/request"
	// "quickstart/services"
	// "quickstart/services/opendata"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// register

func RegisterUser(userAPI *fiber.Ctx) error {
	// var userRegister request.UserRegister
	// if err := userAPI.BodyParser(&userRegister); err != nil {
	// 	logs.WithError(err).Error("cannot parse json")
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }
	// if userRegister.Username == "" || userRegister.Password == "" || userRegister.PasswordConfirm == "" {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }

	// if userRegister.Password != userRegister.PasswordConfirm {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40001, v1.BadRequest, nil)
	// }

	// user := models.OpenDataUser{
	// 	Username:               strings.ToLower(strings.TrimSpace(userRegister.Username)),
	// 	Name:                   userRegister.Name,
	// 	Password:               userRegister.Password,
	// 	MobileNo:               strings.TrimSpace(userRegister.MobileNo),
	// 	OpenDataServiceReq:     true,
	// 	NotificationServiceReq: true,
	// }
	// if userRegister.Organization != "" {
	// 	user.Organization = userRegister.Organization
	// }
	// var err error
	// user, err = opendata.RegisterUser(user)
	// if err != nil {
	// 	if errors.Is(err, opendata.ErrUserDuplicated) {
	// 		return v1.ResponseJsonWithCode(userAPI, 409, 40900, err.Error(), nil)
	// 	}
	// 	logs.WithError(err).Error("cannot register user")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrong, nil)
	// }

	// token, err := opendata.CreateVerifyToken(user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot create verify token")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50001, v1.SomethingWentWrong, nil)
	// }

	// go func() {
	// 	err = mail.SendEmailRegisterUser(user, token.Token)
	// 	if err != nil {
	// 		logs.WithError(err).Error("send mail error")
	// 		//TODO:: add error log
	// 		//services.AddErrorLog(services.PlatformUser, "Send mail register user error", user.Username, 0, &user, nil, time.Now())
	// 	}
	// }()

	return v1.ResponseJsonWithCode(userAPI, 201, 20000, "created", nil)
}

// login
func Login(userAPI *fiber.Ctx) error {

	// var userLogin struct {
	// 	Username string `json:"username"`
	// 	Password string `json:"password"`
	// }
	// if err := userAPI.BodyParser(&userLogin); err != nil {
	// 	logs.WithError(err).Error("cannot parse user login")
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }
	// if userLogin.Username == "" || userLogin.Password == "" {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }
	// user, err := opendata.Login(userLogin.Username, userLogin.Password)
	// if err != nil {
	// 	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
	// 		user = &models.OpenDataUser{
	// 			Username: userLogin.Username,
	// 		}
	// 		// TODO:: open data log
	// 		//go services.AddAccessLog(services.PlatformUser, "Login failed", user)
	// 	}
	// 	logs.WithError(err).Error("user not found")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }
	// if !user.Activated {
	// 	return v1.ResponseJsonWithCode(userAPI, 403, 40300, v1.EmailNotVerified, nil)
	// }
	// if user.Rejected {
	// 	return v1.ResponseJsonWithCode(userAPI, 403, 40302, v1.UserHasBeenRejected, nil)
	// }
	// if !user.Approved {
	// 	return v1.ResponseJsonWithCode(userAPI, 403, 40301, v1.UserNotApproved, nil)
	// }

	// token := opendata.GenerateAccessToken(*user)
	// if token == "" {
	// 	logs.Error("cannot create token")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrong, nil)
	// }
	// err = opendata.SaveAuthenticationToken(token, *user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot save token")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrong, nil)
	// }
	var results = map[string]interface{}{}
	// var results = map[string]interface{}{
	// 	"accessToken": token,
	// 	"permission": map[string]bool{
	// 		"openData":     user.OpenDataServiceApproved,
	// 		"notification": user.NotificationServiceApproved,
	// 	},
	// }
	// // TODO :: Open data log
	// //go services.AddAccessLog(services.PlatformUser, "Login Success", user)
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, results)
}

// logout
func Logout(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user logout")
	// }
	// err = opendata.Logout(v1.GetBearer(userAPI))
	// if err != nil {
	// 	logs.Error("logout:", err)
	// }
	// if user.ID > 0 {
	// 	// TODO :: Open DATA LOGS
	// 	//go services.AddAccessLog(services.PlatformOpenData, "Logout Success", &user)
	// }
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, nil)
}

// permission
func Permission(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }
	// if user.ID > 0 {
	// 	// TODO :: Open DATA LOGS
	// 	//go services.AddAccessLog(services.PlatformOpenData, "Logout Success", &user)
	// }
	var results = map[string]interface{}{}

	// var results = map[string]interface{}{
	// 	"permission": map[string]bool{
	// 		"openData":     user.OpenDataServiceApproved,
	// 		"notification": user.NotificationServiceApproved,
	// 	},
	// }
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, results)
}

func NewPassword(userAPI *fiber.Ctx) error {
	// var userResetPassword struct {
	// 	Password        string `json:"password"`
	// 	PasswordConfirm string `json:"passwordConfirm"`
	// }
	// if err := userAPI.BodyParser(&userResetPassword); err != nil {
	// 	logs.WithError(err).Error("cannot parse user reset password")
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequestTh, nil)
	// }
	// var token = v1.GetBearer(userAPI)
	// if token == "" {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequestTh, nil)
	// }
	// if (userResetPassword.Password == "" || userResetPassword.PasswordConfirm == "") ||
	// 	userResetPassword.Password != userResetPassword.PasswordConfirm {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequestTh, nil)
	// }
	// verifyToken, err := opendata.GetVerifyToken(token, "forgot")
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get verify token")
	// 	return v1.ResponseJsonWithCode(userAPI, 404, 40400, v1.UserNotFoundTh, nil)
	// }
	// user, err := opendata.GetUserByID(verifyToken.OpenDataUserID)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user ")
	// 	return v1.ResponseJsonWithCode(userAPI, 404, 40401, v1.UserNotFoundTh, nil)
	// }
	// // if !v1.ValidPassword(userResetPassword.Password) {
	// // 	return v1.ResponseJsonWithCode(userAPI, 400, 40001, v1.InvalidPassword, nil)
	// // }
	// err = opendata.UpdatePassword(user, userResetPassword.Password)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot update password")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrongTh, nil)
	// }
	// err = opendata.VerifyToken(verifyToken)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot verify reset token")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50001, v1.SomethingWentWrongTh, nil)
	// }
	// // TODO : auto login ?
	// //token := services.GenerateAccessToken(resetToken.User)
	// //if token == "" {
	// //	logs.Error("cannot create token")
	// //	return ResponseJsonWithCode(userAPI, 500, 50000, SomethingWentWrong, nil)
	// //}
	// //err = services.SaveAuthenticationToken(token, resetToken.User)
	// //if err != nil {
	// //	logs.WithError(err).Error("cannot save token")
	// //	return ResponseJsonWithCode(userAPI, 500, 50000, SomethingWentWrong, nil)
	// //}
	// //
	// //var results = map[string]interface{}{
	// //	"accessToken": token,
	// //	"permission": map[string]interface{}{
	// //		"modules": genModuleItems(resetToken.User.RoleID),
	// //	},
	// //}
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, "set password success", nil)
}

func ForgotPassword(userAPI *fiber.Ctx) error {
	// var userForgotPassword request.UserForgotPassword
	// if err := userAPI.BodyParser(&userForgotPassword); err != nil {
	// 	return err
	// }
	// user, err := models.GetOpenDataUserByUsername(userForgotPassword.Username)
	// if err != nil {
	// 	logs.WithError(err).Error("user not found")
	// 	return v1.ResponseJsonWithCode(userAPI, 404, 40400, v1.NotFound, nil)
	// }
	// forgotToken, err := opendata.CreateForgotPasswordToken(user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot create reset token")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50001, v1.SomethingWentWrong, nil)
	// }
	// go func() {
	// 	err := mail.SendEmailOpenDataResetPassword(user, forgotToken.Token)
	// 	if err != nil {
	// 		logs.WithError(err).Error("cannot send mail reset password")
	// 		//TODO:: add error log
	// 		//services.AddErrorLog(services.PlatformUser, "Send mail reset password error", user.Username, 0, &user, nil, time.Now())
	// 	}
	// }()
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.SendMailResetPasswordSuccess, nil)
}

func ChangePassword(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user logout")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }
	// var userResetPassword struct {
	// 	OldPassword     string `json:"oldPassword"`
	// 	Password        string `json:"password"`
	// 	PasswordConfirm string `json:"passwordConfirm"`
	// }
	// if err := userAPI.BodyParser(&userResetPassword); err != nil {
	// 	logs.WithError(err).Error("cannot parse user reset password")
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }
	// if userResetPassword.OldPassword == "" || userResetPassword.PasswordConfirm == "" || userResetPassword.Password != userResetPassword.PasswordConfirm {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40000, v1.BadRequest, nil)
	// }
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userResetPassword.OldPassword))
	// if err != nil {
	// 	logs.Error(err)
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40001, v1.WrongOldPassword, nil)
	// }
	// if !v1.ValidPassword(userResetPassword.Password) {
	// 	return v1.ResponseJsonWithCode(userAPI, 400, 40002, v1.InvalidPassword, nil)
	// }

	// err = opendata.UpdatePassword(user, userResetPassword.Password)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot update pasword")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrong, nil)
	// }
	// // TODO:: auto logout?
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, "change password success", nil)
}

func GetAPIKey(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user logout")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }
	// if !user.OpenDataServiceApproved {
	// 	return v1.ResponseJsonWithCode(userAPI, 403, 40301, v1.Unauthorized, nil)
	// }
	var result = map[string]string{}

	// var result = map[string]string{
	// 	"apiKey":    user.CurrentApiKey,
	// 	"rateLimit": user.CurrentRateLimit,
	// }
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, result)
}

func Profile(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	var results = map[string]interface{}{}
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }

	// var results = map[string]interface{}{
	// 	"email":        user.Username,
	// 	"name":         user.Name,
	// 	"organization": user.Organization,
	// 	"registerDate": user.CreatedAt.Format(v1.DefaultDateTimeFormatResp),
	// 	"permission": map[string]interface{}{
	// 		"openData":     user.OpenDataServiceApproved,
	// 		"notification": user.NotificationServiceApproved,
	// 	},
	// }
	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, results)
}

func DeleteUser(userAPI *fiber.Ctx) error {
	// user, err := getUserAuthentication(userAPI)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot get user")
	// 	return v1.ResponseJsonWithCode(userAPI, 401, 40100, v1.Unauthorized, nil)
	// }
	// // err = models.DeleteUser(user)
	// // if err != nil {
	// // 	logs.WithError(err).Error("cannot delete user")
	// // 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, v1.SomethingWentWrong, nil)
	// // }

	// err = v1.DeleteKongAPIKey(&user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot delete kong api key")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50000, "cannot delete kong api key : "+err.Error(), nil)
	// }
	// err = v1.DeleteKongConsumer(&user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot delete kong consumer")
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50001, "cannot delete kong consumer : "+err.Error(), nil)
	// }
	// err = models.DeleteOpenDataUser(&user)
	// if err != nil {
	// 	logs.WithError(err).Error("cannot delete open data user")
	// 	userSys := userAPI.Locals("user").(models.User)
	// 	go services.AddErrorLog(services.PlatformServiceUser, services.OpenUserDeleted, user.Username, 0, &userSys, nil, time.Now())
	// 	return v1.ResponseJsonWithCode(userAPI, 500, 50002, v1.SomethingWentWrong, nil)
	// }

	// userSys := userAPI.Locals("user").(models.User)
	// go services.AddSystemLog(services.PlatformServiceUser, services.OpenUserDeleted, user.Username, &userSys, nil)

	return v1.ResponseJsonWithCode(userAPI, 200, 20000, v1.Success, nil)
}

// gen token

// func getUserAuthentication(c *fiber.Ctx) (models.OpenDataUser, error) {
// 	token := v1.GetBearer(c)
// 	if token == "" {
// 		return models.OpenDataUser{}, errors.New("authorization not found")
// 	}

// 	user, err := opendata.GetUserByToken(token)
// 	if err != nil {
// 		logs.WithError(err).Error("cannot get user from token: ", token)
// 		return models.OpenDataUser{}, err
// 	}
// 	return user, nil
// }

func generatePrivateKey() (*rsa.PrivateKey, error) {
	// The GenerateKey method takes in a reader that returns random bits, and
	// the number of bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	var privateKeyBytes = x509.MarshalPKCS1PrivateKey(privateKey)
	return privateKeyBytes
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pem file
func generatePublicKey(publicKey *rsa.PublicKey) ([]byte, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return publicKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, keyType, saveFileTo string) error {
	pemKeyBlock := &pem.Block{
		Type:  keyType,
		Bytes: keyBytes,
	}
	pemFile, err := os.Create(saveFileTo)
	if err != nil {
		return err
	}
	err = pem.Encode(pemFile, pemKeyBlock)
	if err != nil {
		return err
	}
	return nil
}

func generateJWTToken(privateKey *rsa.PrivateKey, iss string) (string, error) {
	var claims = jwt.MapClaims{}
	claims["iss"] = iss
	claims["iat"] = time.Now().Unix()
	// TODO:: Exp date?
	claims["exp"] = time.Now().Add(10 * 365 * 24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
