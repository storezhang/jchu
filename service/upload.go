package service

import (
	"github.com/pangum/ft"
)

type Upload struct {
	ft *ft.Client
}

func newUpload(ft *ft.Client) *Upload {
	return &Upload{
		ft: ft,
	}
}
