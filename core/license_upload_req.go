package core

import (
	"os"
	"path/filepath"

	"github.com/drone/envsubst"
)

// LicenseUploadReq 授权请求
type LicenseUploadReq struct {
	Addr   string
	Id     string
	Key    string
	Secret string

	Input      string
	Output     string
	Type       string
	Filename   string
	Result     string
	Enterprise string
	Sheet      string
	Skipped    int
}

func (lur *LicenseUploadReq) RealFilename(name string, code string) (filename string, err error) {
	if name, ne := envsubst.Eval(lur.Filename, lur.env(name, code)); nil != ne {
		err = ne
	} else if LicenseTypeWord == lur.Type {
		filename = filepath.Join(lur.Output, name)
	} else {
		filename = filepath.Join(lur.Input, name)
	}

	return
}

func (lur *LicenseUploadReq) env(name string, code string) func(string) string {
	return func(key string) (value string) {
		switch key {
		case "NAME":
			value = name
		case "CODE":
			value = code
		default:
			value = os.Getenv(key)
		}

		return
	}
}
