package imagehandler

import (
	"content/pkg/domain/response"
	"content/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type ShopHandler struct {
	shopUsecase interfaces.ShopUsecase
}

func NewShopHandler(shopusecase interfaces.ShopUsecase) *ShopHandler {
	return &ShopHandler{shopUsecase: shopusecase}
}

func (sh *ShopHandler) Home(c *gin.Context) {
	allimages, err := sh.shopUsecase.Home()
	if err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to get all images", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	c.JSON(200, allimages)
}

func AddToCart(c *gin.Context){
	
}