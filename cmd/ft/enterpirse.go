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
		Command: cmd.New(`Enterprise`, cmd.Aliases(`ent`, `e`), cmd.Usage(`企业信息`)),

		ft:         ft,
		enterprise: `Enterprise.xlsx`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (e *enterprise) Run(_ *app.Context) (err error) {
	return
}

func (e *enterprise) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`Enterprise`, &e.enterprise, arg.String(e.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewInt(
			`Skipped`, &e.skipped, arg.Int(e.skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`Sheet`, &e.sheet, arg.String(e.sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
