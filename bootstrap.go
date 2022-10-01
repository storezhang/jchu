package main

import (
	"github.com/storezhang/cli/cmd/excel"
	"github.com/storezhang/cli/cmd/word"

	"github.com/pangum/pangu"
)

type (
	bootstrap struct {
		application *pangu.Application
		excel       *excel.Command
		word        *word.Command
	}

	bootstrapIn struct {
		pangu.In

		Application *pangu.Application
		Excel       *excel.Command
		Word        *word.Command
	}
)

func newBootstrap(in bootstrapIn) pangu.Bootstrap {
	return &bootstrap{
		application: in.Application,
		excel:       in.Excel,
		word:        in.Word,
	}
}

func (b *bootstrap) Startup() error {
	return b.application.AddCommands(b.excel, b.word)
}
