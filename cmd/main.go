package main

import (
	"github.com/mereiamangeldin/effective-mobile-test/internal/app"
	"github.com/mereiamangeldin/effective-mobile-test/internal/config"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("initialization of configuration")
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
