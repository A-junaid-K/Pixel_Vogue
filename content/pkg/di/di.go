package di

import (
	"content/pkg/api"
	imagehandler "content/pkg/api/handler"
	"content/pkg/config"
	"content/pkg/db"
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

	// contributorRepository := contributorRepo.NewContributorRepository(gormDB)
	// contributorUsecase := contributorUseCase.NewContributorUseCase(contributorRepository)
	// contributoreHandler := contributorhandler.NewContributorHanlder(contributorUsecase)

	serverHTTP := api.NewServerHTTP(imageHandler)

	return serverHTTP, nil
}
