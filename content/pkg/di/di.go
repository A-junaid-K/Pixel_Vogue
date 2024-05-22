package di

import (
	"content/pkg/api"
	imagehandler "content/pkg/api/handler"
	"content/pkg/config"
	"content/pkg/infrastructure/db"
	"content/pkg/repository"
	"content/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {

	gormDB, err := db.ConnectdB(cfg)
	if err != nil {
		return nil, err
	}

	imageRepository := repository.NewImageRepository(gormDB)
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageHandler := imagehandler.NewImageHandler(imageUsecase)

	shopRepository := repository.NewShopRepository(gormDB)
	shopUsecase := usecase.NewShopUsecase(shopRepository)
	shopHandler := imagehandler.NewShopHandler(shopUsecase)

	// contributorRepository := contributorRepo.NewContributorRepository(gormDB)
	// contributorUsecase := contributorUseCase.NewContributorUseCase(contributorRepository)
	// contributoreHandler := contributorhandler.NewContributorHanlder(contributorUsecase)

	serverHTTP := api.NewServerHTTP(imageHandler,shopHandler)

	return serverHTTP, nil
}
