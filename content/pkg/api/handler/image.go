package imagehandler

import (
	"content/pkg/domain/models"
	"content/pkg/domain/response"
	"content/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ImageHandler struct {
	imageUsecase interfaces.ImageUsecase
}

func NewImageHandler(imageUsecase interfaces.ImageUsecase) *ImageHandler {
	return &ImageHandler{
		imageUsecase: imageUsecase,
	}
}

func (ih *ImageHandler) UploadImage(c *gin.Context) {

	// Verify Contributor

	// Bind
	var body models.Image
	if err := c.Bind(&body); err != nil {
		resp := response.ErrResponse{StatusCode: 500, Response: "Cannot Bind", Error: err.Error()}
		c.JSON(500, resp)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Invalid input", Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	maxsize := int64(50 * 1024 * 1024) // max size 50mb
	if err := c.Request.ParseMultipartForm(maxsize); err != nil{
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to multi parse", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	image,header,err := c.Request.FormFile("image")
	if err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to get form file request", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	if err := ih.imageUsecase.UploadImage(image,*header); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to upload image", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResnpose{StatusCode: 200, Response: "Image Successfully uploaded"}
	c.JSON(200, resp)
}
