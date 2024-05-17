package imagehandler

import (
	"content/pkg/domain/models"
	"content/pkg/domain/response"
	"content/pkg/usecase/interfaces"
	"time"

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

	var body models.ImageDetails

	maxsize := int64(50 * 1024 * 1024) // max size 50mb
	if err := c.Request.ParseMultipartForm(maxsize); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to multi parse", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	contributorId := c.GetInt("id")
	body.ContributorId = contributorId
	body.Size = c.PostForm("size")
	body.Dimension = c.PostForm("dimension")
	body.MoreInformation = c.PostForm("more_information")
	body.Tags = c.PostForm("tags")
	dateTakenstr, err := time.Parse("02-01-2006", c.PostForm("date_taken"))
	if err != nil {
		resp := response.ErrResponse{StatusCode: 500, Response: "Invalid date format", Error: err.Error()}
		c.JSON(400, resp)
		return
	}
	body.DateTaken = dateTakenstr

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

	image, header, err := c.Request.FormFile("image")
	if err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to get form file request", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	if err := ih.imageUsecase.UploadImage(image, *header, body); err != nil {
		resp := response.ErrResponse{StatusCode: 400, Response: "Failed to upload image", Error: err.Error()}
		c.JSON(400, resp)
		return
	}

	resp := response.SuccessResnpose{StatusCode: 200, Response: "Image Successfully uploaded"}
	c.JSON(200, resp)
}

// func (ih *ImageHandler) UploadImage(c *gin.Context, req *pb.UploadImageRequest) {
// 	var body models.Image

// 	body.Details.ContributorId = int(req.ContributorId)

// 	// Bind
// 	if err := c.Bind(&body); err != nil {
// 		resp := response.ErrResponse{StatusCode: 500, Response: "Cannot Bind", Error: err.Error()}
// 		c.JSON(500, resp)
// 		return
// 	}

// }
