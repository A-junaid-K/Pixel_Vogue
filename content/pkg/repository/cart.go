package repository

import (
	"content/pkg/domain/models"
	"content/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &CartRepository{DB: db}
}

func (cr *CartRepository) AddToCart(userid int, cart_item models.Cart) error {

	image := models.Cart{UserID: userid, ImageID: cart_item.ImageID, Quantity: cart_item.Quantity}

	err := cr.DB.Create(&image).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CartRepository) RemoveFromCart(userid int, cart_item models.Cart) error {
	if err := cr.DB.Delete(models.Cart{}).Where("user_id=? AND cart_id=?", userid, cart_item.CartID).Error; err != nil {
		return err
	}

	return nil
}

func (cr *CartRepository) CheckImageInCart(user_id, image_id int) bool {
	var count int64
	cr.DB.Table("carts").Where("user_id=? AND image_id=?", user_id, image_id).Count(&count)

	return count == 0
}
