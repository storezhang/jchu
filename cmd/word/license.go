package word

import (
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/service"

	"github.com/pangum/pangu/app"
)

var _ app.Command = (*license)(nil)

type license struct {
	*cmd.Command

	service *service.License

	license    string
	enterprise string
	output     string
	result     string
	skipped    int
	sheet      string
}

func newLicense(service *service.License) *license {
	return &license{
		Command: cmd.New(`license`, cmd.Aliases(`l`), cmd.Usage(`上传授权协议到大数据中心`)),
		service: service,

		license:    `license.docx`,
		enterprise: `enterprise.xlsx`,
		output:     `license`,
		result:     `result.txt`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (l *license) Run(_ *app.Context) (err error) {
	return l.service.Upload(l.license, l.enterprise, l.output, l.result, l.sheet, l.skipped)
}

func (l *license) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`license`, &l.license, arg.String(l.license),
			arg.Aliases(`l`, `lic`),
			arg.Usage("指定授权`文件`"),
		),
		arg.NewString(
			`enterprise`, &l.enterprise, arg.String(l.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewString(
			`output`, &l.output, arg.String(l.output),
			arg.Aliases(`o`, `out`),
			arg.Usage("指定输出`文件`"),
		),
		arg.NewString(
			`result`, &l.result, arg.String(l.result),
			arg.Aliases(`r`, `res`),
			arg.Usage("指定结果记录`文件`"),
		),
		arg.NewInt(
			`skipped`, &l.skipped, arg.Int(l.skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`sheet`, &l.sheet, arg.String(l.sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
