package handler

import (
	"net/http"

	models "user/pkg/domain/models"
	response "user/pkg/domain/response"
	interfaces "user/pkg/usecase/user/interface"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUsecase interfaces.UserUsecase
}

func NewUserHanlder(userUsecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (uh *UserHandler) UserSignUp(c *gin.Context) {
	var body models.SignUpRequest

	if err := c.Bind(&body); err != nil {
		res := response.ErrResponse{
			StatusCode: 400, Response: "body bind error", Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		resp := response.ErrResponse{
			StatusCode: 400,
			Response:   "invalid input",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := uh.userUsecase.Signup(body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to Signup", Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// resp := response.SuccessResnpose{StatusCode: 303, Response: "Redirect to otp verification"}
	// Redirect to OTP verification
	c.Redirect(303, "/verify-otp")
}

func (uh *UserHandler) VerifyOtp(c *gin.Context) {
	var otp models.VerifyOtp
	if err := c.Bind(&otp); err != nil {
		resp := response.ErrResponse{
			StatusCode: 400, Response: "bind error", Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
	}

	if otp.Otp == 0 {
		resp := response.ErrResponse{
			StatusCode: 400, Response: "otp is empty", Error: "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Verifing Email and Otp
	if err := uh.userUsecase.VerifyOtp(otp.Email, otp.Otp); err != nil {
		resp := response.ErrResponse{
			StatusCode: 400, Response: "Error from OTP varificatoin", Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.SuccessResnpose{StatusCode: 200, Response: "Succefully Signed Up"}
	c.JSON(200, resp)
}

func (uh *UserHandler) UserLogin(c *gin.Context) {
	var loginBody models.LoginRequest

	if err := c.Bind(&loginBody); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	validator := validator.New()
	if err := validator.Struct(&loginBody); err != nil {
		resp := response.ErrResponse{
			StatusCode: http.StatusBadRequest,
			Response:   "Invalid Input",
			Error:      err.Error(),
		}
		c.JSON(400, resp)
		return
	}

	token, err := uh.userUsecase.Login(loginBody)

	if err != nil {
		resp := response.ErrResponse{
			StatusCode: 400,
			Response:   "Login Error",
			Error:      err.Error(),
		}
		c.JSON(400, resp)
		return
	}

	resp := models.LoginResopnse{StatusCode: http.StatusCreated, Token: token}
	c.JSON(201, resp)
}

func (uh *UserHandler) UserProfile(c *gin.Context) {
	userid := c.GetInt("id")
	user, err := uh.userUsecase.GetUserById(userid)
	if err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "User Not Found", Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	userProfile := models.User{
		Email: user.Email,
	}
	c.JSON(200, userProfile)
}

// func (uh *UserHandler) UpdateEmail(c *gin.Context) {
// 	userid := c.GetInt("id")
// 	var body models.UpdateUserProfile
// 	if err := c.Bind(&body); err != nil {
// 		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	validator := validator.New()
// 	if err := validator.Struct(body); err != nil {
// 		res := response.ErrResponse{Response: "Invalid input", Error: err.Error(), StatusCode: 400}
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}
// 	if err := uh.userUsecase.UpdateEmail(body.Email);err != nil{
// 		res := response.ErrResponse{Response: "Failed to Update Email", Error: err.Error(), StatusCode: 400}
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// }
