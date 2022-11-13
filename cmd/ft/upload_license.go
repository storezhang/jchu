package ft

import (
	"github.com/goexl/xiren"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/args"
	"github.com/storezhang/cli/core"
	"github.com/storezhang/cli/service"
)

var _ app.Command = (*upload)(nil)

type upload struct {
	*cmd.Command

	service *service.Upload
	args    uploadArgs
}

func newUpload(service *service.Upload, ft *args.Ft, license *args.License) *upload {
	return &upload{
		Command: cmd.New(`upload`, cmd.Aliases(`u`, `up`), cmd.Usage(`上传`)),

		service: service,
		args: uploadArgs{
			ft:      ft,
			license: license,
		},
	}
}

func (u *upload) Run(_ *app.Context) (err error) {
	if err = xiren.Struct(u.args); nil != err {
		return
	}

	req := new(core.LicenseUploadReq)
	req.Addr = u.args.ft.Addr
	req.Id = u.args.ft.Id
	req.Key = u.args.ft.Key
	req.Secret = u.args.ft.Secret

	req.Type = u.args.license.Type
	req.Input = u.args.license.Input
	req.Output = u.args.license.Output
	req.Filename = u.args.license.Filename
	req.Result = u.args.ft.Result
	req.Enterprise = u.args.license.Enterprise
	req.Sheet = u.args.license.Sheet
	req.Skipped = u.args.license.Skipped

	err = u.service.License(req)

	return
}
