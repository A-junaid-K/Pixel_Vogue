package interfaces

import (
	"mime/multipart"
)

type ImageUsecase interface {
	UploadImage(image multipart.File, head multipart.FileHeader) error
}
