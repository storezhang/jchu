package main

import (
	"cli/command"

	"github.com/pangum/pangu"
)

type (
	bootstrap struct {
		application *pangu.Application
		duplicate   *command.Duplicate
	}

	bootstrapIn struct {
		pangu.In

		Application *pangu.Application
		Duplicate   *command.Duplicate
	}
)

func newBootstrap(in bootstrapIn) pangu.Bootstrap {
	return &bootstrap{
		application: in.Application,
		duplicate:   in.Duplicate,
	}
}

func (b *bootstrap) Setup() error {
	return b.application.AddCommands(b.duplicate)
}
