package license

import (
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/args"
)

var _ app.Command = (*Command)(nil)

type Command struct {
	*cmd.Command

	upload *upload
	args   *args.License
}

func newCommand(args *args.License, upload *upload) *Command {
	return &Command{
		Command: cmd.New("Command").Aliases("l", "lis").Usage("授权协议").Build(),

		args:   args,
		upload: upload,
	}
}

func (c *Command) Subcommands() (commands app.Commands) {
	return app.Commands{
		c.upload,
	}
}

func (c *Command) Arguments() app.Arguments {
	return []app.Argument{
		arg.New[string]("enterprise", &c.args.Enterprise).
			Default(c.args.Enterprise).
			Aliases("e", "ent").
			Usage("指定企业表格`文件`").
			Build(),
		arg.New[string]("type", &c.args.Type).
			Default(c.args.Type).
			Aliases("t", "typ").
			Usage("指定授权文件`类型`").
			Build(),
		arg.New[[]string]("filenames", &c.args.Filenames).
			Default(c.args.Filenames).
			Aliases("fs", "fns").
			Usage("指定授权`文件名`").
			Build(),
		arg.New[string]("input", &c.args.Input).
			Default(c.args.Input).
			Aliases("i", "in").
			Usage("指定授权文件输入`目录`").
			Build(),
		arg.New[string]("output", &c.args.Input).
			Default(c.args.Input).
			Aliases("o", "out").
			Usage("指定授权文件输出`目录`").
			Build(),
		arg.New[int]("skipped", &c.args.Skipped).
			Default(c.args.Skipped).
			Aliases("S", "skip").
			Usage("指定跳过行数").
			Build(),
		arg.New[string]("sheet", &c.args.Sheet).
			Default(c.args.Sheet).
			Aliases("s", "sht").
			Usage("指定企业表格`表名`").
			Build(),
	}
}
