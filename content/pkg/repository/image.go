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

func (ir *ImageRepository) UploadImage(imageUrl string) error {
	var contributor models.Image // fake contributor data
	log.Println("before")
	// if err := ir.DB.Table("images").Where("contributr_id=?", contributor.Id).Create(
	// 	models.Image{
	// 		Image: imageUrl,
	// 	}).Error; err != nil {
	// 	log.Println("db err : ", err)
	// 	return err
	// }

	// if err := ir.DB.Table("images").Where("contributr_id=?", contributor.Id).Set("images", imageUrl).Error; err != nil {
	// 	log.Println("db err : ", err)
	// 	return err
	// }

	log.Println("db after")
	return nil
}
