package command

import (
	"cli/service"

	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
)

type (
	// Duplicate 去重
	Duplicate struct {
		duplicate *service.Duplicate
	}

	duplicateIn struct {
		pangu.In

		Duplicate *service.Duplicate
	}
)

func newDuplicate(in duplicateIn) *Duplicate {
	return &Duplicate{
		duplicate: in.Duplicate,
	}
}

func (d *Duplicate) Run(_ *app.Context) (err error) {
	ins := []string{
		`assert/list.xlsx`,
		`assert/keyword.xlsx`,
	}
	out := `clear.xlsx`
	headers := []string{
		`市州`,
		`区划`,
		`机构名称`,
		`统一社会信用代码`,
		`经营者姓名`,
		`身份证号码`,
		`主体等级`,
		`主体类型`,
		`法定代表人`,
		`行政区划`,
		`成立日期`,
		`注册资金`,
		`行业`,
		`经营状态`,
		`机构地址`,
		`经营范围`,
	}
	if err = d.duplicate.Removal(ins, out, headers); nil != err {
		return
	}

	return
}

func (d *Duplicate) SubCommands() (commands []app.Command) {
	return
}

func (d *Duplicate) Aliases() []string {
	return []string{
		`d`,
	}
}

func (d *Duplicate) Name() string {
	return `duplicate`
}

func (d *Duplicate) Usage() string {
	return ``
}
