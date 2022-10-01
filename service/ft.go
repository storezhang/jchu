package service

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emmansun/gmsm/sm2"
	"github.com/emmansun/gmsm/sm4"
	"github.com/emmansun/gmsm/smx509"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/pangum/http"
	"github.com/pangum/logging"
	"github.com/pangum/pangu"
)

type (
	// Ft 授权协议上传
	Ft struct {
		http       *http.Client
		key        *sm2.PrivateKey
		privateHex string
		publicHex  string
		logger     *logging.Logger
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

func newFt(in ftIn) (ft *Ft, err error) {
	ft = new(Ft)
	ft.http = in.Http
	ft.logger = in.Logger
	if ft.key, err = sm2.GenerateKey(rand.Reader); nil != err {
		return
	}

	if pk, me := smx509.MarshalPKIXPublicKey(ft.key.Public()); nil != me {
		err = me
	} else {
		ft.publicHex = strings.ToUpper(hex.EncodeToString(pk))
	}
	if nil != err {
		return
	}

	if pk, me := smx509.MarshalSM2PrivateKey(ft.key); nil != me {
		err = me
	} else {
		ft.privateHex = strings.ToUpper(hex.EncodeToString(pk))
	}

	return
}

func (f *Ft) request(host string, api string, req any, rsp any) error {
	return f.sendfile(host, api, ``, req, rsp)
}

func (f *Ft) sendfile(host string, api string, file string, req any, rsp any) (err error) {
	hr := f.http.R()
	if form, formErr := gox.StructToForm(req); nil != formErr {
		err = formErr
	} else {
		form[`publicKey`] = f.publicHex
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
	var block cipher.Block
	if key, de := f.key.Decrypt(rand.Reader, []byte(sdr.Key), sm2.NewPlainDecrypterOpts(sm2.C1C2C3)); nil != de {
		err = de
	} else {
		block, err = sm4.NewCipher(key)
	}
	if nil != err {
		return
	}

	var decrypted []byte
	block.Decrypt(decrypted, []byte(sdr.Data))
	err = json.Unmarshal(decrypted, rsp)

	return
}
