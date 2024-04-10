package di

// import (
// 	"user/pkg/api"
// 	handler "user/pkg/api/handler/user"
// 	"user/pkg/config"
// 	"user/pkg/database"
// 	"user/pkg/repository"
// 	"user/pkg/usecase"

// 	"github.com/google/wire"
// )

// func InitApi(cfg config.Config) (*api.ServerHTTP, error) {
// 	wire.Build{
// 		database.ConnectdB,

// 		repository.NewUserRepository,

// 		usecase.NewUserUseCase,

// 		handler.NewUserHanlder,

// 		api.NewServerHTTP,
// 	}
// 	return &api.ServerHTTP{}, nil
// }
