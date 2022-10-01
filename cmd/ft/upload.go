package ft

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*upload)(nil)

type (
	upload struct {
		*cmd.Command

		license *license
	}

	uploadIn struct {
		pangu.In

		License *license
	}
)

func newUpload(in uploadIn) *upload {
	return &upload{
		Command: cmd.New(`upload`, cmd.Aliases(`u`, `up`), cmd.Usage(`文件上传`)),
		license: in.License,
	}
}

func (u *upload) Run(_ *app.Context) (err error) {
	return
}

func (u *upload) Subcommands() (commands []app.Command) {
	return []app.Command{
		u.license,
	}
}
