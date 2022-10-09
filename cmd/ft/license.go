package ft

import (
	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*license)(nil)

type (
	license struct {
		*cmd.Command

		upload   *upload
		template *template
	}

	licenseIn struct {
		pangu.In

		Upload   *upload
		Template *template
	}
)

func newLicense(in licenseIn) *license {
	return &license{
		Command: cmd.New(`license`, cmd.Aliases(`l`), cmd.Usage(`授权协议`)),

		upload:   in.Upload,
		template: in.Template,
	}
}

func (l *license) Run(_ *app.Context) (err error) {
	return
}

func (l *license) Subcommands() (commands []app.Command) {
	return []app.Command{
		l.upload,
		l.template,
	}
}
