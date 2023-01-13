package enterpirse

import (
	"github.com/pangum/ft"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*Command)(nil)

type Command struct {
	*cmd.Command

	ft         *ft.Client
	enterprise string
	skipped    int
	sheet      string
}

func newCommand(ft *ft.Client) *Command {
	return &Command{
		Command: cmd.New("Command").Aliases("ent", "e").Usage("企业信息").Build(),

		ft:         ft,
		enterprise: `Enterprise.xlsx`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("Command", &c.enterprise).
			Default(c.enterprise).
			Aliases("c", "ent").
			Usage("指定企业表格`文件`").
			Build(),
		arg.New("skipped", &c.skipped).
			Default(c.skipped).
			Aliases("S", "skip").
			Usage("指定跳过行数").
			Build(),
		arg.New("sheet", &c.sheet).
			Default(c.sheet).
			Aliases("s", "sht").
			Usage("指定企业表格`表名`").
			Build(),
	}
}
