package service

import (
	"encoding/json"
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/http"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
	"github.com/tjfoc/gmsm/sm4"
)

type (
	// Ft 授权协议上传
	Ft struct {
		http   *http.Client
		logger *logging.Logger
	}

	ftIn struct {
		pangu.In

		Http   *http.Client
		Logger *logging.Logger
	}

	scDataRsp struct {
		// 授权码
		Key string `json:"key"`
		// 签名数据
		SignatureData string `json:"signatureData"`
		// 数据
		Data string `json:"data"`
	}
)

func newFt(in ftIn) *Ft {
	return &Ft{
		http:   in.Http,
		logger: in.Logger,
	}
}

func (f *Ft) request(host string, api string, pk string, req any, rsp any) error {
	return f.sendfile(host, api, pk, ``, req, rsp)
}

func (f *Ft) sendfile(host string, api string, pk string, file string, req any, rsp any) (err error) {
	hr := f.http.R()
	if form, formErr := gox.StructToForm(req); nil != formErr {
		err = formErr
	} else {
		form[`publicKey`] = pk
		form[`token`] = `eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhcHBJZCI6IjE1Njk2MTMyMzE4ODQ4ODYwMTciLCJleHBpcmVzVGltZSI6MTY2NDYzMDY4Mzc5OX0.nDZDWLutKRpbNwq3ltCo4uFQ3V456fxnyeWry8cbre8`
		hr.SetFormData(form)
	}
	if nil != err {
		return
	}

	// 设置上传文件路径
	if `` != file {
		hr.SetFile(`file`, file)
	}

	if raw, reqErr := hr.Post(fmt.Sprintf(`%s%s`, host, api)); nil != reqErr {
		err = reqErr
		f.logger.Error(`发送到省大数据中心出错`, field.String(`api`, api), field.Error(err))
	} else if raw.IsError() {
		f.logger.Warn(`发送到省大数据中心出错`, field.String(`api`, api), field.String(`raw`, raw.String()))
	} else {
		err = f.decrypt(raw.Body(), rsp)
	}

	return
}

func (f *Ft) decrypt(raw []byte, rsp any) (err error) {
	sdr := new(scDataRsp)
	if err = json.Unmarshal(raw, sdr); nil != err {
		return
	}

	// 解密
	if block, smErr := sm4.NewCipher([]byte(sdr.Key)); nil != smErr {
		err = smErr
	} else {
		var decrypted []byte
		block.Decrypt(decrypted, []byte(sdr.Data))
		err = json.Unmarshal(decrypted, rsp)
	}

	return
}
