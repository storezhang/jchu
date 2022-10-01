package ft

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goexl/gfx"
	"github.com/nguyenthenguyen/docx"
	"github.com/pangum/pangu/arg"
	"github.com/pangum/pangu/cmd"
	"github.com/storezhang/cli/core"
	"github.com/storezhang/cli/service"
	"github.com/xuri/excelize/v2"

	"github.com/pangum/pangu/app"
)

var _ app.Command = (*license)(nil)

type license struct {
	*cmd.Command

	ft *service.Ft

	license    string
	enterprise string
	output     string
	skipped    int
	sheet      string
}

func newLicense(ft *service.Ft) *license {
	return &license{
		Command: cmd.New(`license`, cmd.Aliases(`l`), cmd.Usage(`上传授权协议到大数据中心`)),
		ft:      ft,

		license:    `license.docx`,
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
			if success, uploadErr := l.upload(l.license, resultFile, columns...); nil != uploadErr || !success {
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

func (l *license) upload(license string, result *os.File, columns ...string) (success bool, err error) {
	req := new(core.FtLicenseUploadReq)
	req.Name = columns[0]
	req.Code = columns[1]
	req.AuthorizedInfos = columns[2]
	req.AuthorizedCode = columns[3]
	req.PlatformId = columns[4]
	req.Count = columns[5]
	req.LoanStage = columns[6]
	req.FileType = columns[7]
	req.CaSupplier = columns[8]
	req.ValidateUrl = columns[9]
	req.AuthorizedStartTime = columns[10]
	req.AuthorizedEndTime = columns[11]

	var doc *docx.Docx
	if file, readErr := docx.ReadDocxFile(license); nil != readErr {
		err = readErr
	} else {
		doc = file.Editable()
		defer func() {
			err = file.Close()
		}()
	}
	if nil != err {
		return
	}

	if err = doc.Replace(`[Name]`, req.Name, -1); nil != err {
		return
	}
	if err = doc.Replace(`[Code]`, req.Code, -1); nil != err {
		return
	}

	// 转换成PDF格式的文件
	realFile := filepath.Join(l.output, fmt.Sprintf(`%s.docx`, strings.TrimSpace(req.Name)))
	if err = doc.WriteToFile(realFile); nil != err {
		return
	}

	if rsp, ue := l.ft.Upload(host, pk, realFile, req); nil != ue {
		err = ue
	} else {
		_, err = result.WriteString(fmt.Sprintf("%s\t\t%s\t\t%s", req.Name, req.Code, rsp.LicenseId))
	}

	return
}
