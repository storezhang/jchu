package service

import (
	"github.com/pangum/ft"
	"github.com/pangum/logging"
)

type Upload struct {
	ft     *ft.Client
	logger *logging.Logger
}

func newUpload(ft *ft.Client, logger *logging.Logger) *Upload {
	return &Upload{
		ft:     ft,
		logger: logger,
	}
}
