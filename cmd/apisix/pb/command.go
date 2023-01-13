package pb

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

		args   *pbArgs
		upload *upload
	}

	commandIn struct {
		pangu.In

		Args   *pbArgs
		Upload *upload
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New("pb").Aliases("protobuf", "proto", "p").Usage("Protobuf协议").Build(),

		args:   in.Args,
		upload: in.Upload,
	}
}

func (c *Command) Subcommands() (commands app.Commands) {
	return app.Commands{
		c.upload,
	}
}

func (c *Command) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("id", &c.args.Id).
			Aliases("i", "identify").
			Usage("指定协议`编号`，可以是任意字符，建议尽量选择有意义的编号").
			Build(),
	}
}
