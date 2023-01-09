package apisix

import (
	"context"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/pangum/apisix"
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/args"
	"github.com/storezhang/cli/core"
)

var _ app.Command = (*pb)(nil)

type (
	pb struct {
		*cmd.Command
		simaqian.Logger

		creator *apisix.Creator
		server  *args.TokenServer

		id          string
		filename    string
		description string
	}

	pbIn struct {
		pangu.In

		Creator *apisix.Creator
		Server  *args.TokenServer
		Logger  core.Logger
	}
)

func newPb(in pbIn) *pb {
	return &pb{
		Command: cmd.New("pb").Aliases("protobuf", "proto", "p").Usage("Protobuf协议").Build(),
		Logger:  in.Logger,

		creator: in.Creator,
		server:  in.Server,
	}
}

func (p *pb) Run(_ *app.Context) (err error) {
	client := p.creator.Create(p.server.Addr, p.server.Token)
	fields := gox.Fields[any]{
		field.New("filename", p.filename),
		field.New("id", p.id),
		field.New("description", p.description),
	}
	if rsp, ue := client.UploadDescriptor(context.Background(), p.filename, p.id, p.description); nil != ue {
		err = ue
		p.Warn("上传协议文件出错", fields.Connect(field.Error(ue))...)
	} else {
		p.Warn("上传协议文件成功", fields.Connect(field.New("rsp", rsp))...)
	}

	return
}

func (p *pb) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("id", &p.id).
			Default(p.id).
			Aliases("id", "identify").
			Usage("指定协议`编号`，可以是任意字符，建议尽量选择有意义的编号").
			Build(),
		arg.New("filename", &p.filename).
			Default(p.filename).
			Aliases("f", "fn").
			Usage("指定协议`文件`").
			Build(),
		arg.New("description", &p.description).
			Default(p.description).
			Aliases("d", "desc").
			Usage("指定协议`描述信息`").
			Build(),
	}
}
