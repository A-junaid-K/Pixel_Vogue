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

func (ir *ImageRepository) UploadImage(imageUrl string) error {
	var contributor models.Image // fake contributor data

	if err := ir.DB.Table("images").Where("contributr_id=?", contributor.Id).Set("image", imageUrl).Error; err != nil {
		return err
	}

	return nil
}
