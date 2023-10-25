package app

import (
	"github.com/mereiamangeldin/effective-mobile-test/internal/config"
	"github.com/mereiamangeldin/effective-mobile-test/internal/handler"
	"github.com/mereiamangeldin/effective-mobile-test/internal/repository/pg"
	"github.com/mereiamangeldin/effective-mobile-test/internal/service"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/httpserver"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/logging"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	logger := logging.GetLogger()
	logger.Info("starting an application")

	logger.Info("connecting to database")
	db, err := pg.New(
		logger,
		pg.WithHost(cfg.DB.Host),
		pg.WithPort(cfg.DB.Port),
		pg.WithDBName(cfg.DB.DBName),
		pg.WithUsername(cfg.DB.Username),
		pg.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		logger.Infof("connection to DB err: %s", err.Error())
		return err
	}

	logger.Info("creating service part")
	srvs := service.NewManager(db, cfg, logger)
	logger.Info("creating handlers")
	hndlr := handler.New(srvs, logger)
	logger.Info("connecting to server")
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	logger.Info("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		logger.Infof("signal received: %s", s.String())
	case err = <-server.Notify():
		logger.Infof("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		logger.Infof("server shutdown err: %s", err)
	}

	return nil
}
