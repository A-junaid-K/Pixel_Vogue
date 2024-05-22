package interfaces

type ShopUsecase interface {
	// Home()([]models.Image,error)
	Home() ([]string, error)
	AddToCart()
	RemoveFromCart()
	AddToWishlist()
	RemoveFromWishlist()
	Search()
}
