package interfaces

import "content/pkg/domain/models"

type CartUsecase interface {
	AddToCart(userid int, cart_item models.Cart) error
}
