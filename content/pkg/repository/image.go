package repository

import (
	"content/pkg/domain/models"
	"content/pkg/repository/interfaces"
	"log"

	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) interfaces.ImageRepository {
	return &ImageRepository{DB: db}
}

func (ir *ImageRepository) UploadImage(imageUrl string, imagedetails models.ImageDetails) error {

	image := &models.Image{Image: imageUrl}
	log.Println("iamge repository.")
	if err := ir.DB.Create(image).Error; err != nil {
		return err
	}

	imageDetails := &models.ImageDetails{
		ContributorId:   imagedetails.ContributorId,
		ImageID:         image.Id,
		Size:            imagedetails.Size,
		Dimension:      imagedetails.Dimension,
		DateTaken:       imagedetails.DateTaken,
		MoreInformation: imagedetails.MoreInformation,
		Tags:            imagedetails.Tags,
	}

	if err := ir.DB.Create(imageDetails).Error; err != nil {
		return err
	}

	// if err := ir.DB.Table("images").Where("contributr_id=?", contributorId).Set("images", imageUrl).Error; err != nil {
	// 	log.Println("db err : ", err)
	// 	return err
	// }

	return nil
}
