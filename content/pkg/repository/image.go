package repository

import (
	"content/pkg/domain/models"
	"content/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) interfaces.ImageRepository {
	return &ImageRepository{DB: db}
}

func (ir *ImageRepository) UploadImage(imageUrl string, imagedetails models.ImageDetails) error {

	imageDetails := &models.ImageDetails{
		ContributorId: imagedetails.ContributorId,
		// ImageID:         image.Id,
		Size:            imagedetails.Size,
		Dimension:       imagedetails.Dimension,
		DateTaken:       imagedetails.DateTaken,
		MoreInformation: imagedetails.MoreInformation,
		Tags:            imagedetails.Tags,
	}

	image := &models.Image{
		Image:   imageUrl,
		Details: *imageDetails,
	}

	if err := ir.DB.Create(image).Error; err != nil {
		return err
	}

	// if err := ir.DB.Create(imageDetails).Error; err != nil {
	// 	return err
	// }

	// if err := ir.DB.Table("images").Where("contributr_id=?", contributorId).Set("images", imageUrl).Error; err != nil {
	// 	log.Println("db err : ", err)
	// 	return err
	// }

	return nil
}

func (ir *ImageRepository) GetImage(id int) (models.Image, error) {
	var empty, image models.Image

	err := ir.DB.Table("image_details").Where("contributor_id=?", id).First(&image).Error
	if err != nil {
		return empty, err
	}

	return image, nil
}
