package service

import (
	"github.com/mereiamangeldin/effective-mobile-test/internal/config"
	"github.com/mereiamangeldin/effective-mobile-test/internal/repository"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/logging"
)

type Manager struct {
	Repository repository.IRepository
	Config     *config.Config
	logger     logging.Logger
}

func NewManager(repository repository.IRepository, config *config.Config, logger logging.Logger) *Manager {
	return &Manager{Repository: repository, Config: config, logger: logger}
}
