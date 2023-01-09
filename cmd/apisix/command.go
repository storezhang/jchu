package apisix

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/args"
)

var _ app.Command = (*Command)(nil)

type (
	// Command 命令
	Command struct {
		*cmd.Command

		args *args.TokenServer
		pb   *pb
	}

	commandIn struct {
		pangu.In

		Args *args.TokenServer
		Pb   *pb
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New("apisix").Usage("Apisix网关命令").Aliases("as").Build(),

		args: in.Args,
		pb:   in.Pb,
	}
}

func (c *Command) Subcommands() (commands app.Commands) {
	return app.Commands{
		c.pb,
	}
}

func (c *Command) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("endpoint", &c.args.Addr).
			Default(c.args.Addr).
			Aliases("e", "ep").
			Usage("指定服务器`端点`").
			Build(),
		arg.New("key", &c.args.Token).
			Default(c.args.Token).
			Aliases("k", "ak").
			Usage("指定应用`通信密钥`").
			Build(),
	}
}
