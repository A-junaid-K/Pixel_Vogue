package api

import (
	"fmt"
	"net/http"
	userhandler "user/pkg/api/handler/user"
	contributorhandler "user/pkg/api/handler/contributor"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHanlder *userhandler.UserHandler , contributorHandler *contributorhandler.ContributorHandler) *ServerHTTP {
	router := gin.New()

	userAuthRoute := router.Group("/user")
	{
		userAuthRoute.POST("/signup", userHanlder.UserSignUp)
		userAuthRoute.POST("/verify-otp", userHanlder.VerifyOtp)
		userAuthRoute.POST("/login", userHanlder.UserLogin)
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
