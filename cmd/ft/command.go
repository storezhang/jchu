package ft

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

		args       *args.Ft
		license    *license
		enterprise *enterprise
	}

	commandIn struct {
		pangu.In

		Args       *args.Ft
		License    *license
		Enterprise *enterprise
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New("ft").Usage("52号文相关命令").Aliases("52").Build(),

		args:       in.Args,
		license:    in.License,
		enterprise: in.Enterprise,
	}
}

func (c *Command) Subcommands() (commands app.Commands) {
	return app.Commands{
		c.license,
		c.enterprise,
	}
}

func (c *Command) Arguments() app.Arguments {
	return app.Arguments{
		arg.New[string]("id", &c.args.Id).
			Default(c.args.Id).
			Aliases("i", "identify").
			Usage("指定应用`编号`").
			Build(),
		arg.New[string]("key", &c.args.Key).
			Default(c.args.Key).
			Aliases("k", "ak").
			Usage("指定应用`用户名`").
			Build(),
		arg.New[string]("secret", &c.args.Secret).
			Default(c.args.Secret).
			Aliases("s", "sk").
			Usage("指定应用`密码`").
			Build(),
		arg.New[string]("addr", &c.args.Secret).
			Default(c.args.Secret).
			Aliases("a", "address").
			Usage("指定接口`地址`").
			Build(),
		arg.New[string]("result", &c.args.Result).
			Default(c.args.Result).
			Aliases("r", "res").
			Usage("指定结果记录`文件`").
			Build(),
	}
}
