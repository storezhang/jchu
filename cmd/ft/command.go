package ft

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*Command)(nil)

type (
	// Command 命令
	Command struct {
		*cmd.Command

		upload *upload
	}

	commandIn struct {
		pangu.In

		Upload *upload
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New(`fifty-two`, cmd.Usage(`52号文相关命令`), cmd.Aliases(`5`, `52`, `ft`)),
		upload:  in.Upload,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Subcommands() (commands []app.Command) {
	return []app.Command{
		c.upload,
	}
}

func (c *Command) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`host`, &host, arg.String(host),
			arg.Aliases(`hst`),
			arg.Usage("指定接口`地址`"),
		),
	}
}
