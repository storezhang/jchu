package ft

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goexl/ft"
	"github.com/nguyenthenguyen/docx"
)

func (u *upload) action(result *os.File, columns ...string) (success bool, err error) {
	req := new(ft.LicenseUploadReq)
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
	if file, dfe := docx.ReadDocxFromFS(`template.docx`, licenseFS); nil != dfe {
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

	if err = doc.Replace(`[Name]`, req.Name, -1); nil != err {
		return
	}
	if err = doc.Replace(`[Code]`, req.Code, -1); nil != err {
		return
	}

	// 转换成PDF格式的文件
	realFile := filepath.Join(u.output, fmt.Sprintf(`%s.docx`, strings.TrimSpace(req.Name)))
	if err = doc.WriteToFile(realFile); nil != err {
		return
	}

	if rsp, ue := u.ft.Upload(realFile, req, ft.Addr(addr), ft.App(id, key, secret)); nil != ue {
		err = ue
	} else {
		_, err = result.WriteString(fmt.Sprintf("%s\t\t%s\t\t%s", req.Name, req.Code, rsp.LicenseId))
	}

	return
}
