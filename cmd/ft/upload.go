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

var _ app.Command = (*upload)(nil)

type upload struct {
	*cmd.Command

	ft         *ft.Client
	license    string
	enterprise string
	output     string
	skipped    int
	sheet      string
}

func newUpload(ft *ft.Client) *upload {
	return &upload{
		Command: cmd.New(`license`, cmd.Aliases(`u`, `up`), cmd.Usage(`文件上传`)),

		ft:         ft,
		license:    `license.docx`,
		enterprise: `enterprise.xlsx`,
		output:     `license`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}

func (u *upload) Run(_ *app.Context) (err error) {
	if _, exists := gfx.Exists(u.output); !exists {
		err = os.MkdirAll(u.output, os.ModePerm)
	}
	if nil != err {
		return
	}

	var excel *excelize.File
	if excel, err = excelize.OpenFile(u.enterprise); nil != err {
		return
	}

	var rows *excelize.Rows
	if rows, err = excel.Rows(u.sheet); nil != err {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// 跳过N行
	for i := 0; i < u.skipped; i++ {
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
			if success, ue := u.action(u.license, resultFile, columns...); nil != ue || !success {
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}

	return
}

func (u *upload) Args() []app.Arg {
	return []app.Arg{
		arg.NewString(
			`license`, &u.license, arg.String(u.license),
			arg.Aliases(`u`, `lic`),
			arg.Usage("指定授权`文件`"),
		),
		arg.NewString(
			`enterprise`, &u.enterprise, arg.String(u.enterprise),
			arg.Aliases(`e`, `ent`),
			arg.Usage("指定企业表格`文件`"),
		),
		arg.NewString(
			`result`, &result, arg.String(result),
			arg.Aliases(`r`, `res`),
			arg.Usage("指定结果记录`文件`"),
		),
		arg.NewInt(
			`skipped`, &u.skipped, arg.Int(u.skipped),
			arg.Aliases(`S`, `skip`),
			arg.Usage("指定跳过行数"),
		),
		arg.NewString(
			`sheet`, &u.sheet, arg.String(u.sheet),
			arg.Aliases(`s`, `sht`),
			arg.Usage("指定企业表格`表名`"),
		),
	}
}
