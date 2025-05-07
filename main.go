package main

import (
	"github.com/carinhanha/gin/config"
	"github.com/carinhanha/gin/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errf("config initializarion error %v", err)
		return
	}

	router.Initialize()
}