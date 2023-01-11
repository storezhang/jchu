package pb

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

var _ app.Command = (*upload)(nil)

type (
	upload struct {
		*cmd.Command
		simaqian.Logger

		creator *apisix.Creator
		args    uploadArgs

		filename    string
		description string
	}

	uploadArgs struct {
		server *args.TokenServer
		pb     *args.Pb
	}

	uploadIn struct {
		pangu.In

		Creator *apisix.Creator
		Server  *args.TokenServer
		Pb      *args.Pb
		Logger  core.Logger
	}
)

func newUpload(in uploadIn) *upload {
	return &upload{
		Command: cmd.New("upload").Aliases("up", "u").Usage("上传Protobuf协议").Build(),
		Logger:  in.Logger,

		creator: in.Creator,
		args: uploadArgs{
			server: in.Server,
			pb:     in.Pb,
		},
	}
}

func (p *upload) Run(_ *app.Context) (err error) {
	client := p.creator.Create(p.args.server.Addr, p.args.server.Token)
	fields := gox.Fields[any]{
		field.New("filename", p.filename),
		field.New("id", p.args.pb.Id),
		field.New("description", p.description),
	}
	if rsp, ue := client.UploadDescriptor(context.Background(), p.filename, p.args.pb.Id, p.description); nil != ue {
		err = ue
		p.Warn("上传协议文件出错", fields.Connect(field.Error(ue))...)
	} else {
		p.Debug("上传协议文件成功", fields.Connect(field.New("rsp", rsp))...)
	}

	return
}

func (p *upload) Arguments() app.Arguments {
	return app.Arguments{
		arg.New("filename", &p.filename).
			Aliases("f", "fn").
			Usage("指定协议`文件`").
			Build(),
		arg.New("description", &p.description).
			Aliases("d", "desc").
			Usage("指定协议`描述信息`").
			Build(),
	}
}
