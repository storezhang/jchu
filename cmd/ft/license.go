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
		Command: cmd.New("license").Aliases("l", "lis").Usage("授权协议").Build(),

		args:   args,
		upload: upload,
	}
}

func (l *license) Subcommands() (commands app.Commands) {
	return app.Commands{
		l.upload,
	}
}

func (l *license) Arguments() app.Arguments {
	return []app.Argument{
		arg.New[string]("enterprise", &l.args.Enterprise).
			Default(l.args.Enterprise).
			Aliases("e", "ent").
			Usage("指定企业表格`文件`").
			Build(),
		arg.New[string]("type", &l.args.Type).
			Default(l.args.Type).
			Aliases("t", "typ").
			Usage("指定授权文件`类型`").
			Build(),
		arg.New[[]string]("filenames", &l.args.Filenames).
			Default(l.args.Filenames).
			Aliases("fs", "fns").
			Usage("指定授权`文件名`").
			Build(),
		arg.New[string]("input", &l.args.Input).
			Default(l.args.Input).
			Aliases("i", "in").
			Usage("指定授权文件输入`目录`").
			Build(),
		arg.New[string]("output", &l.args.Input).
			Default(l.args.Input).
			Aliases("o", "out").
			Usage("指定授权文件输出`目录`").
			Build(),
		arg.New[int]("skipped", &l.args.Skipped).
			Default(l.args.Skipped).
			Aliases("S", "skip").
			Usage("指定跳过行数").
			Build(),
		arg.New[string]("sheet", &l.args.Sheet).
			Default(l.args.Sheet).
			Aliases("s", "sht").
			Usage("指定企业表格`表名`").
			Build(),
	}
}
