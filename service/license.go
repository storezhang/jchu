package service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/gfx"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/nguyenthenguyen/docx"
	"github.com/pangum/http"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"github.com/storezhang/cli/core"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/xuri/excelize/v2"
)

type (
	// License 授权协议上传
	License struct {
		http   *http.Client
		logger *logging.Logger
	}

	licenseIn struct {
		pangu.In

		Http   *http.Client
		Logger *logging.Logger
	}
)

func newLicense(in licenseIn) *License {
	return &License{
		http:   in.Http,
		logger: in.Logger,
	}
}

func (l *License) Upload(license, enterprise, output, result, sheet string, skipped int) (err error) {
	if _, exists := gfx.Exists(output); exists {
		err = os.MkdirAll(output, os.ModePerm)
	}
	if nil != err {
		return
	}

	var excel *excelize.File
	if excel, err = excelize.OpenFile(enterprise); nil != err {
		return
	}

	var rows *excelize.Rows
	if rows, err = excel.Rows(sheet); nil != err {
		return
	}
	defer func() {
		err = rows.Close()
	}()

	// 跳过N行
	for i := 0; i < skipped; i++ {
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
			if success, uploadErr := l.upload(license, resultFile, columns...); nil != uploadErr || !success {
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}

	return
}

func (l *License) upload(license string, output *os.File, columns ...string) (success bool, err error) {
	name := columns[0]
	code := columns[1]
	authorizedInfos := columns[2]
	authorizedOrgCode := columns[3]
	platformId := columns[4]
	count := columns[5]
	loanStage := columns[6]
	licenseFileType := columns[7]
	caSupplier := columns[8]
	validateUrl := columns[9]
	authorizedStartTime := columns[10]
	authorizedEndTime := columns[11]

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

	if err = doc.Replace(`[Name]`, name, -1); nil != err {
		return
	}
	if err = doc.Replace(`[Code]`, code, -1); nil != err {
		return
	}

	// 转换成PDF格式的文件
	// gex.Exec(`pandoc`, gex.Args(``))
	realFile := fmt.Sprintf(`license/%s.docx`, name)
	if err = doc.WriteToFile(realFile); nil != err {
		return
	}

	var hash string
	if data, readErr := ioutil.ReadFile(realFile); nil != readErr {
		return
	} else {
		sm := sm3.New()
		sm.Write(data)
		hash = strings.ToUpper(hex.EncodeToString(sm.Sum(nil)))
	}

	url := fmt.Sprintf(`https://202.61.91.57:8092/api/creditInquiry/uploadLicense`)
	data := map[string]string{
		`publicKey`:           `041F60E35FBB8EEC21BBBAA3F8CBF78D104F23FF3F32A12A30E4B9B2A77CA4DCC52495676B9DCF277BA5B85BA7057D1E06F5FB9B997F072DB967521F64852BD2BB`,
		`token`:               `eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhcHBJZCI6IjE1Njk2MTMyMzE4ODQ4ODYwMTciLCJleHBpcmVzVGltZSI6MTY2NDYzMDY4Mzc5OX0.nDZDWLutKRpbNwq3ltCo4uFQ3V456fxnyeWry8cbre8`,
		`enterpriseName`:      name,
		`uniscId`:             code,
		`authorizedInfos`:     authorizedInfos,
		`authorizedOrgsCode`:  authorizedOrgCode,
		`platformId`:          platformId,
		`count`:               count,
		`loanStage`:           loanStage,
		`licenseFileType`:     licenseFileType,
		`caSupplier`:          caSupplier,
		`validateUrl`:         validateUrl,
		`authorizedStartTime`: authorizedStartTime,
		`authorizedEndTime`:   authorizedEndTime,
		`hashCode`:            hash,
	}

	var raw *resty.Response
	if raw, err = l.http.R().SetFormData(data).SetFile(`file`, realFile).Post(url); nil != err {
		return
	}

	rsp := new(core.SCDataResponse)
	if err = json.Unmarshal(raw.Body(), rsp); nil != err {
		return
	}

	fields := gox.Fields{
		field.String(`name`, name),
		field.String(`code`, code),
		field.String(`raw`, raw.String()),
	}
	if rsp.Check() {
		_, err = output.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", code, rsp.Key, rsp.SignatureData, rsp.Data))
		success = true
		l.logger.Warn(`上传协议成功`, fields...)
	} else {
		l.logger.Warn(`上传协议失败`, fields...)
	}

	return
}
