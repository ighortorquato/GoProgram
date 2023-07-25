package main

import (
	"github.com/ighortorquato/GoProgram.git/config"
	r "github.com/ighortorquato/GoProgram.git/router"
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
