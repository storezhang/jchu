package service

import (
	"encoding/hex"
	"io/ioutil"
	"strings"

	"github.com/storezhang/cli/core"
	"github.com/tjfoc/gmsm/sm3"
)

func (f *Ft) Upload(host, license string, req *core.FtLicenseUploadReq) (rsp *core.FtLicenseUploadRsp, err error) {
	if data, readErr := ioutil.ReadFile(license); nil != readErr {
		err = readErr
	} else {
		sm := sm3.New()
		sm.Write(data)
		req.HashCode = strings.ToUpper(hex.EncodeToString(sm.Sum(nil)))
	}
	if nil != err {
		return
	}

	rsp = new(core.FtLicenseUploadRsp)
	err = f.sendfile(host, `/api/creditInquiry/uploadLicense`, license, req, rsp)

	return
}
