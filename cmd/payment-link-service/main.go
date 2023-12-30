package main

import (
	"log"

	"github.com/rifqoi/xendok-service/internal/config"
	"github.com/rifqoi/xendok-service/internal/logger"
)

func main() {
	args := config.ProcessArgs()

	err := config.Init(args)
	if err != nil {
		log.Panicf("Failed to initialize config: %v", err)
	}

	logger.Get()

}
