package api

import (
	imagehandler "content/pkg/api/handler"
	"content/pkg/api/middleware"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(imageHandler *imagehandler.ImageHandler) *ServerHTTP {

	router := gin.New()

	//Config Cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}

	router.Use(cors.Default())

	image := router.Group("/image")
	{
		image.POST("/upload", middleware.ContributorAuth, imageHandler.UploadImage)
	}

	return &ServerHTTP{engine: router}
}

func (r *ServerHTTP) Start(port string) error {
	fmt.Println("server listening in port : ", port)
	if err := http.ListenAndServe(port, r.engine); err != nil {
		return err
	}
	return nil
}
