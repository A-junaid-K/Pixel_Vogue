package main

import (
	"content/pkg/config"
	"content/pkg/di"
	"log"
)

func main() {
	cfg := config.InitConfig()

	server, err := di.InitApi(cfg)
	if err != nil {
		log.Println("Api initialization err : ", err)
	}

	if err := server.Start(cfg.AppPort); err != nil {
		log.Println("server error : ", err)
	}

}
