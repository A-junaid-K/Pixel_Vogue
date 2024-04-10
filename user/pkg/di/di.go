package di

import (
	"user/pkg/api"
	userhandler "user/pkg/api/handler/user"
	contributorhandler "user/pkg/api/handler/contributor"
	"user/pkg/config"
	"user/pkg/database"
	contributor"user/pkg/repository/contributor"
	user"user/pkg/repository/user"
	userUsecase"user/pkg/usecase/user"
	contributorUseCase"user/pkg/usecase/contributor"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {

	gormDB, err := database.ConnectdB(cfg)
	if err != nil {
		return nil, err
	}

	userRepository := user.NewUserRepository(gormDB)
	userUsecase := userUsecase.NewUserUseCase(userRepository)
	userHandler := userhandler.NewUserHanlder(userUsecase)

	contributorRepository := contributor.NewContributorRepository(gormDB)
	contributorUsecase := contributorUseCase.NewContributorUseCase(contributorRepository)
	contributoreHandler := contributorhandler.NewContributorHanlder(contributorUsecase)

	serverHTTP := api.NewServerHTTP(userHandler,contributoreHandler)

	return serverHTTP, nil
}
