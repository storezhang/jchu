package apisix

import (
	"github.com/pangum/ft"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*pb)(nil)

type pb struct {
	*cmd.Command

	ft        *ft.Client
	filename  string
	filenames []string
}

func newPb(ft *ft.Client) *pb {
	return &pb{
		Command: cmd.New("pb").Aliases("protobuf", "proto", "p").Usage("Protobuf协议").Build(),

		ft: ft,
	}
}

func (p *pb) Run(_ *app.Context) (err error) {
	return
}

func (p *pb) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("filename", &p.filename).
			Default(p.filename).
			Aliases("f", "fn").
			Usage("指定协议`文件`").
			Build(),
		arg.New("filenames", &p.filenames).
			Default(p.filenames).
			Aliases("fs", "fns").
			Usage("指定协议`文件`列表").
			Build(),
	}
}
