package handler

import (
	"github.com/mereiamangeldin/effective-mobile-test/internal/service"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/logging"
)

type Handler struct {
	srvs   service.IService
	logger logging.Logger
}

func New(srvs service.IService, logger logging.Logger) *Handler {
	return &Handler{
		srvs:   srvs,
		logger: logger,
	}
}
