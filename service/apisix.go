package service

import (
	"github.com/pangum/http"
	"github.com/storezhang/cli/core"
)

type Apisix struct {
	http   *http.Client
	logger core.Logger
}

func newApisix(http *http.Client, logger core.Logger) *Apisix {
	return &Apisix{
		http:   http,
		logger: logger,
	}
}
