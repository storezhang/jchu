package ft

import (
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/args"
)

var _ app.Command = (*license)(nil)

type license struct {
	*cmd.Command

	upload *upload
	args   *args.License
}

func newLicense(args *args.License, upload *upload) *license {
	return &license{
		Command: cmd.New("license", cmd.Aliases("lis", "l"), cmd.Usage("授权协议")),

		args:   args,
		upload: upload,
	}
}

func (l *license) Subcommands() (commands []app.Command) {
	return []app.Command{
		l.upload,
	}
}

func (l *license) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`enterprise`, &l.args.Enterprise, arg.String(l.args.Enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewString(
			`type`, &l.args.Type, arg.String(l.args.Type),
			arg.Aliases(`t`, `typ`),
			arg.Usage("指定授权文件`类型`"),
		),
		arg.NewString(
			`filename`, &l.args.Filename, arg.String(l.args.Filename),
			arg.Aliases(`f`, `fn`),
			arg.Usage("指定授权`文件名`"),
		),
		arg.NewString(
			`input`, &l.args.Input, arg.String(l.args.Input),
			arg.Aliases(`i`, `in`),
			arg.Usage("指定授权文件输入`目录`"),
		),
		arg.NewString(
			`output`, &l.args.Output, arg.String(l.args.Output),
			arg.Aliases(`o`, `out`),
			arg.Usage("指定授权文件输出`目录`"),
		),
		arg.NewInt(
			`skipped`, &l.args.Skipped, arg.Int(l.args.Skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`sheet`, &l.args.Sheet, arg.String(l.args.Sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
