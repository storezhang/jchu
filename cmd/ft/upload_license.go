package ft

import (
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/service"
)

var _ app.Command = (*upload)(nil)

type (
	upload struct {
		*cmd.Command

		service *service.Upload
		args    uploadArgs
	}

	uploadArgs struct {
		command *args
		license *licenseArgs
	}
)

func newUpload(service *service.Upload, args *args, license *licenseArgs) *upload {
	return &upload{
		Command: cmd.New(`upload`, cmd.Aliases(`u`, `up`), cmd.Usage(`上传`)),

		service: service,
		args: uploadArgs{
			command: args,
			license: license,
		},
	}
}

func (u *upload) Run(_ *app.Context) error {
	req := new(service.LicenseReq)
	req.Addr = u.args.command.addr
	req.Id = u.args.command.id
	req.Key = u.args.command.key
	req.Secret = u.args.command.secret

	req.Output = u.args.license.output
	req.Result = u.args.command.result
	req.Enterprise = u.args.license.enterprise
	req.Sheet = u.args.license.sheet
	req.Skipped = u.args.license.skipped

	return u.service.License(req)
}
