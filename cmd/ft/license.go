package ft

import (
	"os"
	"time"

	"github.com/goexl/gfx"
	"github.com/pangum/ft"
	"github.com/pangum/pangu/app"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/xuri/excelize/v2"
)

var _ app.Command = (*license)(nil)

type license struct {
	*cmd.Command

	ft         *ft.Client
	enterprise string
	output     string
	skipped    int
	sheet      string
}

func newLicense(ft *ft.Client) *license {
	return &license{
		Command: cmd.New(`license`, cmd.Aliases(`l`), cmd.Usage(`授权协议`)),

		ft:         ft,
		enterprise: `enterprise.xlsx`,
		output:     `license`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (l *license) Run(_ *app.Context) (err error) {
	if _, exists := gfx.Exists(l.output); !exists {
		err = os.MkdirAll(l.output, os.ModePerm)
	}
	if nil != err {
		return
	}

	var excel *excelize.File
	if excel, err = excelize.OpenFile(l.enterprise); nil != err {
		return
	}

	var rows *excelize.Rows
	if rows, err = excel.Rows(l.sheet); nil != err {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// 跳过N行
	for i := 0; i < l.skipped; i++ {
		rows.Next()
	}

	var resultFile *os.File
	if resultFile, err = os.OpenFile(result, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm); nil != err {
		return
	}
	defer func() {
		_ = resultFile.Close()
	}()

	var columns []string
	for rows.Next() {
		if columns, err = rows.Columns(); nil != err {
			continue
		}

		for {
			if success, ue := l.upload(resultFile, columns...); nil != ue || !success {
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}

	return
}

func (l *license) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`enterprise`, &l.enterprise, arg.String(l.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewString(
			`result`, &result, arg.String(result),
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
