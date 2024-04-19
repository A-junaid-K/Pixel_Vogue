package api

import (
	"fmt"
	"net/http"
	contributorhandler "user/pkg/api/handler/contributor"
	userhandler "user/pkg/api/handler/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *userhandler.UserHandler, contributorHandler *contributorhandler.ContributorHandler) *ServerHTTP {
	router := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}

	router.Use(cors.Default())

	userAuthRoute := router.Group("/user")
	{
		userAuthRoute.POST("/signup", userHandler.UserSignUp)
		userAuthRoute.POST("/verify-otp", userHandler.VerifyOtp)
		userAuthRoute.POST("/login", userHandler.UserLogin)

		userAuthRoute.POST("/profile", userHandler.UserProfile)
	}

	contributorAuthRoute := router.Group("/contributor")
	{
		contributorAuthRoute.POST("/register", contributorHandler.ContributorRegister)
		contributorAuthRoute.POST("/verify-otp", contributorHandler.VerifyOtp)
		contributorAuthRoute.POST("/login", contributorHandler.ContributorLogin)
	}

	return &ServerHTTP{engine: router}
}

func (r *ServerHTTP) Start(port string) error {
	fmt.Println("server listening in port: ", port)
	if err := http.ListenAndServe(port, r.engine); err != nil {
		return err
	}
	return nil
}
