package ft

import (
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/cmd"
)

var _ app.Command = (*template)(nil)

type template struct {
	*cmd.Command
}

func newTemplate() *template {
	return &template{
		Command: cmd.New(`template`, cmd.Aliases(`t`, `tpl`), cmd.Usage(`文件模板`)),
	}
}

func (t *template) Run(_ *app.Context) (err error) {
	return
}
