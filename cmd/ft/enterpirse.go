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
		Command: cmd.New(`enterprise`, cmd.Aliases(`ent`, `e`), cmd.Usage(`企业信息`)),

		ft:         ft,
		enterprise: `enterprise.xlsx`,
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
			`enterprise`, &e.enterprise, arg.String(e.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewString(
			`result`, &result, arg.String(result),
			arg.Aliases(`r`, `res`),
			arg.Usage("指定结果记录`文件`"),
		),
		arg.NewInt(
			`skipped`, &e.skipped, arg.Int(e.skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`sheet`, &e.sheet, arg.String(e.sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
