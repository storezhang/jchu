package excel

import (
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/service"

	"github.com/pangum/pangu"
	"github.com/pangum/pangu/app"
)

var _ app.Command = (*merge)(nil)

type (
	merge struct {
		*cmd.Command

		duplicate *service.Duplicate
	}

	mergeIn struct {
		pangu.In

		Duplicate *service.Duplicate
	}
)

func newMerge(in mergeIn) *merge {
	return &merge{
		Command: cmd.New("duplicate").Aliases("d", "dup").Usage(`去重`).Build(),

		duplicate: in.Duplicate,
	}
}

func (m *merge) Run(_ *app.Context) (err error) {
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
	if err = m.duplicate.Removal(ins, out, headers); nil != err {
		return
	}

	return
}
