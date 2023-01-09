package service

import (
	"github.com/pangum/apisix"
	"github.com/pangum/http"
	"github.com/storezhang/cli/core"
)

type Apisix struct {
	http    *http.Client
	logger  core.Logger
	creator *apisix.Creator
}

func newApisix(http *http.Client, logger core.Logger, creator *apisix.Creator) *Apisix {
	return &Apisix{
		http:    http,
		logger:  logger,
		creator: creator,
	}
}
