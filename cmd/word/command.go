package word

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
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
		Command: cmd.New(`doc`, cmd.Usage(`处理Office Word文件`), cmd.Aliases(`w`, `word`, `docx`)),
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
