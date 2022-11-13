package ft

import (
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*license)(nil)

type license struct {
	*cmd.Command

	upload *upload
	args   *licenseArgs
}

func newLicense(args *licenseArgs, upload *upload) *license {
	return &license{
		Command: cmd.New("license", cmd.Aliases("lis", "l"), cmd.Usage("授权协议")),

		args:   args,
		upload: upload,
	}
}

func (l *license) Run(_ *app.Context) (err error) {
	return
}

func (l *license) Subcommands() (commands []app.Command) {
	return []app.Command{
		l.upload,
	}
}

func (l *license) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`enterprise`, &l.args.enterprise, arg.String(l.args.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewInt(
			`skipped`, &l.args.skipped, arg.Int(l.args.skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`sheet`, &l.args.sheet, arg.String(l.args.sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
