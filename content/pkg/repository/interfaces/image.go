package interfaces

type ImageRepository interface {
	UploadImage(imageUrl, contributorId string) error
}
