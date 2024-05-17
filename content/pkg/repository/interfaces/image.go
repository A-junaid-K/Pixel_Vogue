package interfaces

import "content/pkg/domain/models"

type ImageRepository interface {
	UploadImage(imageUrl string, body models.ImageDetails) error
}
