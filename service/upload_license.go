package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goexl/ft"
	"github.com/goexl/gfx"
	"github.com/nguyenthenguyen/docx"
	"github.com/storezhang/cli/asset"
	"github.com/xuri/excelize/v2"
)

func (u *Upload) License(req *LicenseReq) (err error) {
	if _, exists := gfx.Exists(req.Output); !exists {
		err = os.MkdirAll(req.Output, os.ModePerm)
	}
	if nil != err {
		return
	}

	var excel *excelize.File
	if excel, err = excelize.OpenFile(req.Enterprise); nil != err {
		return
	}

	var rows *excelize.Rows
	if rows, err = excel.Rows(req.Sheet); nil != err {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// 跳过N行
	for i := 0; i < req.Skipped; i++ {
		rows.Next()
	}

	var resultFile *os.File
	if resultFile, err = os.OpenFile(req.Result, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm); nil != err {
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
			if success, ue := u.license(req, resultFile, columns); nil != ue || !success {
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}

	return
}

func (u *Upload) license(req *LicenseReq, result *os.File, columns []string) (success bool, err error) {
	lur := new(ft.LicenseUploadReq)
	lur.Name = columns[0]
	lur.Code = columns[1]
	lur.AuthorizedInfos = columns[2]
	lur.AuthorizedCode = columns[3]
	lur.PlatformId = columns[4]
	lur.Count = columns[5]
	lur.LoanStage = columns[6]
	lur.FileType = columns[7]
	lur.CaSupplier = columns[8]
	lur.ValidateUrl = columns[9]
	lur.AuthorizedStartTime = columns[10]
	lur.AuthorizedEndTime = columns[11]

	var doc *docx.Docx
	if file, dfe := docx.ReadDocxFromFS(`template.docx`, asset.License); nil != dfe {
		err = dfe
	} else {
		doc = file.Editable()
		defer func() {
			err = file.Close()
		}()
	}
	if nil != err {
		return
	}

	if err = doc.Replace(`[Name]`, lur.Name, -1); nil != err {
		return
	}
	if err = doc.Replace(`[Code]`, lur.Code, -1); nil != err {
		return
	}

	// 转换成PDF格式的文件
	realFile := filepath.Join(req.Output, fmt.Sprintf(`%s.docx`, strings.TrimSpace(lur.Name)))
	if err = doc.WriteToFile(realFile); nil != err {
		return
	}

	if rsp, ue := u.ft.Upload(realFile, lur, ft.Addr(req.Addr), ft.App(req.Id, req.Key, req.Secret)); nil != ue {
		err = ue
	} else {
		_, err = result.WriteString(fmt.Sprintf("%s\t\t%s\t\t%s", lur.Name, lur.Code, rsp.LicenseId))
	}

	return
}
