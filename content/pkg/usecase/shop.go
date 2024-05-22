package usecase

import (
	"content/pkg/repository/interfaces"
	usecase "content/pkg/usecase/interfaces"
)

type ShopUsecase struct {
	ShopRepo interfaces.ShopRepository
}

func NewShopUsecase(shoprepo interfaces.ShopRepository) usecase.ShopUsecase {
	return &ShopUsecase{ShopRepo: shoprepo}
}

func (su *ShopUsecase) Home() ([]string, error) {
	allimages, err := su.ShopRepo.Home()
	if err != nil {
		return allimages, err
	}
	return allimages, nil
}

func (su *ShopUsecase) AddToCart() {}

func (su *ShopUsecase) RemoveFromCart() {}

func (su *ShopUsecase) AddToWishlist() {}

func (su *ShopUsecase) RemoveFromWishlist() {}

func (su *ShopUsecase) Search() {}
