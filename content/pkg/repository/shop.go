package repository

import (
	"content/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type ShopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) interfaces.ShopRepository {
	return &ShopRepository{DB: db}
}

func (sr *ShopRepository) Home() ([]string, error) {

	var allimages []string

	// var allimages []image
	if err := sr.DB.Table("images").Select([]string{"image"}).Find(&allimages).Error; err != nil {
		return allimages, err
	}
	return allimages, nil
}

func (sr *ShopRepository) AddToCart() {}

func (sr *ShopRepository) RemoveFromCart() {}
