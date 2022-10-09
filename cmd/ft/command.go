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

		license *license
	}

	commandIn struct {
		pangu.In

		License *license
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New(`fifty-two`, cmd.Usage(`52号文相关命令`), cmd.Aliases(`5`, `52`, `ft`)),
		license: in.License,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Subcommands() (commands []app.Command) {
	return []app.Command{
		c.license,
	}
}

func (c *Command) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`id`, &addr, arg.String(addr),
			arg.Aliases(`identify`, `ai`, `i`),
			arg.Usage("指定应用`编号`"),
		),
		arg.NewString(
			`key`, &addr, arg.String(addr),
			arg.Aliases(`key`, `ak`, `k`),
			arg.Usage("指定应用`用户名`"),
		),
		arg.NewString(
			`secret`, &addr, arg.String(addr),
			arg.Aliases(`as`, `sk`, `s`),
			arg.Usage("指定接口`地址`"),
		),
		arg.NewString(
			`addr`, &addr, arg.String(addr),
			arg.Aliases(`address`, `a`),
			arg.Usage("指定接口`地址`"),
		),
	}
}
