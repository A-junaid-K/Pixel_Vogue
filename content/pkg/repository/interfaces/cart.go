package interfaces

import "content/pkg/domain/models"

type CartRepository interface {
	AddToCart(userid int,cart_item models.Cart) error
	RemoveFromCart(userid int,cart_item models.Cart) error
	CheckImageInCart(user_id, image_id int) bool
}
