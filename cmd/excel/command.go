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
		Command: cmd.New(`excel`, cmd.Usage(`处理Office Excel表格文件`), cmd.Aliases(`e`, `xlsx`)),
		merge:   in.Merge,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Subcommands() (commands []app.Command) {
	return []app.Command{
		c.merge,
	}
}

func (c *Command) Args() []app.Arg {
	return []app.Arg{
		arg.NewStrings(
			`inputs`, &c.inputs,
			arg.Aliases(`i`, `ins`),
			arg.Usage("指定输入`文件列表`"),
			arg.Required(),
		),
		arg.NewString(
			`output`, &c.output,
			arg.Aliases(`o`, `out`),
			arg.Usage("指定输出`文件`"),
			arg.Required(),
		),
	}
}
