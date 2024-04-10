package main

import (
	"fmt"
	"user/pkg/config"
	"user/pkg/di"
)

func main() {

	cfg := config.InitConfig()

	server, err := di.InitApi(cfg)

	if err != nil {
		fmt.Println("InitApi error: ", err)
	}

	if err := server.Start(cfg.AppPort); err != nil {
		fmt.Println("server error: ", err)
	}

}
