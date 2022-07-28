package main

import (
	"go-wire-architecture/pkg/config"
	"go-wire-architecture/pkg/di"
	"log"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error loading config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("error starting server: ", diErr)
	} else {
		server.Start()
	}
}
