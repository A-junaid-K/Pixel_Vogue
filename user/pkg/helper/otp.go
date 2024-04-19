package helper

import (
	"math/rand"
	"net/http"
	"net/smtp"

	"user/pkg/config"
	"user/pkg/database"
	models "user/pkg/domain/models"
	response "user/pkg/domain/response"

	"github.com/gin-gonic/gin"
)

func GenerateOTP() int {
	return rand.Intn(9000) + 1000
}

func SendOtp(otp, email string) error {
	cfg := config.GetConfig()

	auth := smtp.PlainAuth("", cfg.Email, cfg.EmailPassword, "smtp.gmail.com")
	to := []string{email}
	message := "Subject: Otp verification\nyour verification otp is " + otp
	return smtp.SendMail("smtp.gmail.com:587", auth, cfg.Email, to, []byte(message))
}
	
type OtpVerifiaction struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}

func VerifyOtp(c *gin.Context) {
	// otp and geting from user
	var otp OtpVerifiaction
	if err := c.Bind(&otp); err != nil {
		resp := response.ErrResponse{
			StatusCode: 422,
			Response:   "failed to parse request body. Please ensure it's valid JSON",
			Error:      err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", otp.Email).First(&user).Error; err != nil {
		resp := response.ErrResponse{
			StatusCode: 404,
			Response:   "User No Found",
			Error:      err.Error(),
		}
		c.JSON(http.StatusNotFound, resp)
		return
	}

	// checking the user already validate or not.
	// if user.Validate {
	// 	resp := utilities.Response{
	// 		StatusCode: 409,
	// 		Err:        "User is already verified. Please log in directly",
	// 		Data:       nil,
	// 	}
	// 	utilities.ResponseResult(c, resp)
	// 	return
	// }

	// checking the otp correct or not.
	if otp.Otp != user.Otp {
		resp := response.ErrResponse{
			StatusCode: 400,
			Response:   "Invalid OTP entered. Please check your OTP and try again.",
			Error:      "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// update validate = true
	database.DB.Model(&models.User{}).Where("id = ?", user.Id).Update("validate", true)

	resp := response.SuccessResnpose{
		StatusCode: 200,
		Response:   "Successfully otp varified",
	}
	c.JSON(http.StatusOK, resp)
}
