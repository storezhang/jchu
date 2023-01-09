package main

import (
	"github.com/storezhang/cli/cmd/apisix"
	"github.com/storezhang/cli/cmd/excel"
	"github.com/storezhang/cli/cmd/ft"

	"github.com/pangum/pangu"
)

type (
	bootstrap struct {
		application *pangu.Application
		excel       *excel.Command
		ft          *ft.Command
		apisix      *apisix.Command
	}

	bootstrapIn struct {
		pangu.In

		Application *pangu.Application
		Excel       *excel.Command
		Ft          *ft.Command
		Apisix      *apisix.Command
	}
)

func newBootstrap(in bootstrapIn) pangu.Bootstrap {
	return &bootstrap{
		application: in.Application,
		excel:       in.Excel,
		ft:          in.Ft,
		apisix:      in.Apisix,
	}
}

func (b *bootstrap) Startup() error {
	return b.application.AddCommands(b.excel, b.ft, b.apisix)
}
