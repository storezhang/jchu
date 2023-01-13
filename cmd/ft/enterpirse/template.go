package enterpirse

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
		Command: cmd.New("template").Aliases(`t`, `tpl`).Usage(`文件模板`).Build(),
	}
}

func (t *template) Run(_ *app.Context) (err error) {
	return
}
