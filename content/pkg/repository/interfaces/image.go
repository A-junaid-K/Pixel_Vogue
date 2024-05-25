package interfaces

import "content/pkg/domain/models"

type ImageRepository interface {
	UploadImage(imageUrl string, body models.ImageDetails) error
	GetImage(id int)(models.Image, error)
	CheckImageExists(id int)bool
}
