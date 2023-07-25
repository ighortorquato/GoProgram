package main

import (
	r "github.com/ighortorquato/GoProgram.git/router"
	"github.com/ighortorquato/GoProgram.git/schemas/config"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize router
	r.Initialize()
}
