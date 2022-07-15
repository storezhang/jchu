package excel

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
)

type (
	// Command 命令
	Command struct {
		merge *merge
	}

	commandIn struct {
		pangu.In

		Merge *merge
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		merge: in.Merge,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) SubCommands() (commands []app.Command) {
	return []app.Command{
		c.merge,
	}
}

func (c *Command) Aliases() []string {
	return []string{
		`e`,
	}
}

func (c *Command) Name() string {
	return `excel`
}

func (c *Command) Usage() string {
	return ``
}
