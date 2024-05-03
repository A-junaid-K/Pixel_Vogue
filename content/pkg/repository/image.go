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

func (ir *ImageRepository) UploadImage(imageUrl, contributorId string) error {

	if err := ir.DB.Table("images").Where("contributr_id=?", contributorId).Create(
		models.Image{
			Image: imageUrl,
		}).Error; err != nil {
		return err
	}

	if err := ir.DB.Table("image_details").Where("contributr_id=?", contributorId).Create(
		models.Image{
			Image: imageUrl,
		}).Error; err != nil {
		return err
	}

	if err := ir.DB.Table("images").Where("contributr_id=?", contributorId).Set("images", imageUrl).Error; err != nil {
		log.Println("db err : ", err)
		return err
	}

	return nil
}
