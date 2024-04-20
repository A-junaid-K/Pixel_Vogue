package interfaces

type ImageRepository interface {
	UploadImage(imageUrl string) error
}
