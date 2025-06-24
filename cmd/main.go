package main

import (
	"projectIO/config"
	"projectIO/internal/app"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)

}
