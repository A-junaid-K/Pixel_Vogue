package di

import (
	"user/pkg/api"
	"user/pkg/config"
	"user/pkg/database"

	userhandler "user/pkg/api/handler/user"
	userRepo "user/pkg/repository/user"
	userUsecase "user/pkg/usecase/user"

	contributorhandler "user/pkg/api/handler/contributor"
	contributorRepo "user/pkg/repository/contributor"
	contributorUseCase "user/pkg/usecase/contributor"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {

	gormDB, err := database.ConnectdB(cfg)
	if err != nil {
		return nil, err
	}

	userRepository := userRepo.NewUserRepository(gormDB)
	userUsecase := userUsecase.NewUserUseCase(userRepository)
	userHandler := userhandler.NewUserHanlder(userUsecase)

	contributorRepository := contributorRepo.NewContributorRepository(gormDB)
	contributorUsecase := contributorUseCase.NewContributorUseCase(contributorRepository)
	contributoreHandler := contributorhandler.NewContributorHanlder(contributorUsecase)

	serverHTTP := api.NewServerHTTP(userHandler, contributoreHandler)

	return serverHTTP, nil
}
