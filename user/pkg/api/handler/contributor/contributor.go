package contributorHandler

import (
	"log"
	"net/http"
	models "user/pkg/domain/models"
	response "user/pkg/domain/response"
	interfaces "user/pkg/usecase/contributor/interface"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

type ContributorHandler struct {
	contributorUsecase interfaces.ContributorUsecase
}

func NewContributorHanlder(contributorUsecase interfaces.ContributorUsecase) *ContributorHandler {
	return &ContributorHandler{
		contributorUsecase: contributorUsecase,
	}
}

func (uh *ContributorHandler) ContributorRegister(c *gin.Context) {
	var body models.Contributor

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

	if err := uh.contributorUsecase.Register(body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to Register", Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// resp := response.SuccessResnpose{StatusCode: 303, Response: "Redirect to otp verification"}
	// Redirect to OTP verification
	c.Redirect(303, "/verify-otp")

}

func (uh *ContributorHandler) VerifyOtp(c *gin.Context) {
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
	if err := uh.contributorUsecase.VerifyOtp(otp.Email, otp.Otp); err != nil {
		resp := response.ErrResponse{
			StatusCode: 400, Response: "Error from OTP varificatoin", Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := response.SuccessResnpose{StatusCode: 200, Response: "Succefully registered Contributor"}
	c.JSON(200, resp)
}

func (uh *ContributorHandler) ContributorLogin(c *gin.Context) {
	var loginBody models.LoginRequest

	if err := c.Bind(&loginBody); err != nil {
		res := response.ErrResponse{Response: "Binding Error", Error: err.Error(), StatusCode: 400}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := uh.contributorUsecase.Login(loginBody)

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

func (uh *ContributorHandler) CallUploadImage(c *gin.Context) {

	const port = ":8000"

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect server: ", err)
		return
	}

	defer conn.Close()

	client := proto.pb.NewUploadImage

}
