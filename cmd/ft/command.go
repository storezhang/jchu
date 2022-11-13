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

		args       *args
		license    *license
		enterprise *enterprise
	}

	commandIn struct {
		pangu.In

		Args       *args
		License    *license
		Enterprise *enterprise
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New("ft", cmd.Usage("52号文相关命令"), cmd.Aliases("52")),

		args:       in.Args,
		license:    in.License,
		enterprise: in.Enterprise,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Subcommands() (commands []app.Command) {
	return []app.Command{
		c.license,
		c.enterprise,
	}
}

func (c *Command) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			"id", &c.args.id, arg.String(c.args.id),
			arg.Aliases("identify", "i"),
			arg.Usage("指定应用`编号`"),
		),
		arg.NewString(
			"key", &c.args.key, arg.String(c.args.key),
			arg.Aliases("ak", "k"),
			arg.Usage("指定应用`用户名`"),
		),
		arg.NewString(
			"secret", &c.args.secret, arg.String(c.args.secret),
			arg.Aliases("sk", "s"),
			arg.Usage("指定接口`地址`"),
		),
		arg.NewString(
			"addr", &c.args.addr, arg.String(c.args.addr),
			arg.Aliases("address", "a"),
			arg.Usage("指定接口`地址`"),
		),
		arg.NewString(
			`result`, &c.args.result, arg.String(c.args.result),
			arg.Aliases(`r`, `res`),
			arg.Usage("指定结果记录`文件`"),
		),
	}
}
