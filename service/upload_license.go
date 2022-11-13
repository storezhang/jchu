package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/ft"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/nguyenthenguyen/docx"
	"github.com/storezhang/cli/asset"
	"github.com/storezhang/cli/core"
	"github.com/xuri/excelize/v2"
)

func (u *Upload) License(req *core.LicenseUploadReq) (err error) {
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

		success := true
		name := columns[0]
		code := columns[1]
		for count := 0; count < 10; count++ {
			if success, err = u.license(name, code, req, resultFile, columns[2:]); nil != err || !success {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}

		subdirectory := "成功"
		fields := gox.Fields{
			field.String("企业名称", name),
			field.String("统一代码", code),
		}
		if !success {
			subdirectory = "失败"
			u.logger.Warn("授权协议上传失败", fields.Connect(field.Error(err))...)
		} else {
			u.logger.Info("授权协议上传成功", fields...)
		}
		if path, pe := req.RealFilename(name, code); nil != pe {
			err = pe
		} else {
			err = gfx.Rename(path, filepath.Join(filepath.Dir(path), subdirectory, gfx.Name(path, gfx.Ext(filepath.Ext(path)))))
		}
	}

	return
}

func (u *Upload) license(
	name string, code string,
	req *core.LicenseUploadReq,
	result *os.File,
	contents []string,
) (success bool, err error) {
	lur := new(ft.LicenseUploadReq)
	lur.Name = name
	lur.Code = code
	lur.AuthorizedInfos = contents[0]
	lur.AuthorizedCode = contents[1]
	lur.PlatformId = contents[2]
	lur.Count = contents[3]
	lur.LoanStage = contents[4]
	lur.FileType = contents[5]
	lur.CaSupplier = contents[6]
	lur.ValidateUrl = contents[7]
	lur.AuthorizedStartTime = contents[8]
	lur.AuthorizedEndTime = contents[9]

	if file, re := u.realFile(req, lur); nil != re {
		err = re
	} else if rsp, ue := u.ft.Upload(file, lur, ft.Addr(req.Addr), ft.App(req.Id, req.Key, req.Secret)); nil != ue {
		err = ue
	} else {
		_, err = result.WriteString(fmt.Sprintf("%s\t\t%s\t\t%s\n", name, code, rsp.LicenseId))
	}

	return
}

func (u *Upload) realFile(req *core.LicenseUploadReq, lur *ft.LicenseUploadReq) (filename string, err error) {
	switch req.Type {
	case core.LicenseTypeWord:
		filename, err = u.fromWord(req, lur)
	case core.LicenseTypeDirect:
		filename, err = u.fromDirect(req, lur)
	}
	if _, exists := gfx.Exists(filename); !exists {
		err = exc.NewFields("文件不存在", field.String("企业名称", lur.Name), field.String("统一代码", lur.Code))
	}

	return
}

func (u *Upload) fromDirect(req *core.LicenseUploadReq, lur *ft.LicenseUploadReq) (string, error) {
	return req.RealFilename(lur.Name, lur.Code)
}

func (u *Upload) fromWord(req *core.LicenseUploadReq, lur *ft.LicenseUploadReq) (filename string, err error) {
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
	if filename, err = req.RealFilename(lur.Name, lur.Code); nil == err {
		err = doc.WriteToFile(filename)
	}

	return
}
