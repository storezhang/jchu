package excel

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

		merge *merge

		inputs []string
		output string
	}

	commandIn struct {
		pangu.In

		Merge *merge
	}
)

func newCommand(in commandIn) *Command {
	return &Command{
		Command: cmd.New("excel").Usage("处理Office Excel表格文件").Aliases("e", "xlsx").Build(),
		merge:   in.Merge,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Subcommands() (commands app.Commands) {
	return app.Commands{
		c.merge,
	}
}

func (c *Command) Arguments() app.Arguments {
	return app.Arguments{
		arg.New[[]string]("inputs", &c.inputs).
			Aliases("i", "ins").
			Usage("指定输入`文件列表`").
			Build(),
		arg.New[string]("output", &c.output).
			Aliases("o", "out").
			Required().
			Usage("指定输出`文件`").
			Build(),
	}
}
