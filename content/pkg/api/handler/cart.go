package imagehandler

import (
	"content/pkg/domain/models"
	"content/pkg/domain/response"
	"content/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CartHandler struct {
	CartUsecase interfaces.CartUsecase
}

func NewCartHandler(cartusecase interfaces.CartUsecase) *CartHandler {
	return &CartHandler{CartUsecase: cartusecase}
}

func (ch *CartHandler) AddToCart(c *gin.Context) {
	var body models.Cart
	userid := c.GetInt("id")
	if err := c.Bind(&body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Bind Error", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to Validate", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	if err := ch.CartUsecase.AddToCart(userid,body); err != nil {
		resp := response.ErrResponse{StatusCode: 500, Response: "Failed to Add the image in cart", Error: err.Error()}
		c.JSON(500, resp)
		return
	}

	resp := response.SuccessResnpose{StatusCode: 200, Response: "Successfully added cart"}
	c.JSON(200, resp)
}
