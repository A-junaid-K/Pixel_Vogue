package interfaces

import (
	"content/pkg/domain/models"
	"mime/multipart"
)

type ImageUsecase interface {
	UploadImage(image multipart.File, head multipart.FileHeader, body models.ImageDetails) error
}
