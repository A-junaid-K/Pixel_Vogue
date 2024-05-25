package interfaces

type ShopRepository interface {
	// Home() ([]models.Image, error)
	Home() ([]string, error)
}
