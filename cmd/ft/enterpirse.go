package ft

import (
	"github.com/pangum/ft"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*enterprise)(nil)

type enterprise struct {
	*cmd.Command

	ft         *ft.Client
	enterprise string
	skipped    int
	sheet      string
}

func newEnterprise(ft *ft.Client) *enterprise {
	return &enterprise{
		Command: cmd.New("enterprise").Aliases("ent", "e").Usage("企业信息").Build(),

		ft:         ft,
		enterprise: `Enterprise.xlsx`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (e *enterprise) Run(_ *app.Context) (err error) {
	return
}

func (e *enterprise) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("enterprise", &e.enterprise).
			Default(e.enterprise).
			Aliases("e", "ent").
			Usage("指定企业表格`文件`").
			Build(),
		arg.New("skipped", &e.skipped).
			Default(e.skipped).
			Aliases("S", "skip").
			Usage("指定跳过行数").
			Build(),
		arg.New("sheet", &e.sheet).
			Default(e.sheet).
			Aliases("s", "sht").
			Usage("指定企业表格`表名`").
			Build(),
	}
}
