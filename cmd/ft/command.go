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
		Command: cmd.New("ft", cmd.Usage("52号文相关命令"), cmd.Aliases("52")),

		args:       in.Args,
		license:    in.License,
		enterprise: in.Enterprise,
	}
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
			"id", &c.args.Id, arg.String(c.args.Id),
			arg.Aliases("identify", "i"),
			arg.Usage("指定应用`编号`"),
		),
		arg.NewString(
			"key", &c.args.Key, arg.String(c.args.Key),
			arg.Aliases("ak", "k"),
			arg.Usage("指定应用`用户名`"),
		),
		arg.NewString(
			"secret", &c.args.Secret, arg.String(c.args.Secret),
			arg.Aliases("sk", "s"),
			arg.Usage("指定接口`地址`"),
		),
		arg.NewString(
			"addr", &c.args.Addr, arg.String(c.args.Addr),
			arg.Aliases("address", "a"),
			arg.Usage("指定接口`地址`"),
		),
		arg.NewString(
			`result`, &c.args.Result, arg.String(c.args.Result),
			arg.Aliases(`r`, `res`),
			arg.Usage("指定结果记录`文件`"),
		),
	}
}
