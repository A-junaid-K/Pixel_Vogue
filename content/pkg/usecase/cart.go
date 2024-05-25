package usecase

import (
	"content/pkg/domain/models"
	"content/pkg/repository/interfaces"
	usecase "content/pkg/usecase/interfaces"
	"errors"
)

type CartUsecase struct {
	CartRepo  interfaces.CartRepository
	ImageRepo interfaces.ImageRepository
}

func NewCartUsecase(cartrepo interfaces.CartRepository) usecase.CartUsecase {
	return &CartUsecase{CartRepo: cartrepo}
}

func (cu *CartUsecase) AddToCart(userid int, cart_item models.Cart) error {
	// invalid image id
	if ok := cu.ImageRepo.CheckImageExists(cart_item.ImageID); !ok {
		return errors.New("there is no image in this ID")
	}

	// already in cart
	if ok := cu.CartRepo.CheckImageInCart(userid, cart_item.CartID); !ok {
		return errors.New("this image already exist in cart")
	}

	if err := cu.CartRepo.AddToCart(userid, cart_item); err != nil {
		return err
	}

	return nil
}
