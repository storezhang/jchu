package service

import (
	"github.com/pangum/ft"
)

type (
	Upload struct {
		ft *ft.Client
	}

	LicenseReq struct {
		Addr   string
		Id     string
		Key    string
		Secret string

		Output     string
		Result     string
		Enterprise string
		Sheet      string
		Skipped    int
	}
)

func newUpload(ft *ft.Client) *Upload {
	return &Upload{
		ft: ft,
	}
}
