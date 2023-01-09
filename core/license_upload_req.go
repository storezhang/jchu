package core

import (
	"os"
	"path/filepath"

	"github.com/drone/envsubst"
	"github.com/goexl/exc"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/field"
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
	Filenames  []string
	Result     string
	Enterprise string
	Sheet      string
	Skipped    int
}

func (lur *LicenseUploadReq) RealFilename(name string, code string) (filename string, err error) {
	for _, _filename := range lur.Filenames {
		if filename, err = lur.filename(_filename, name, code); nil == err {
			break
		}
	}
	if "" != filename {
		err = nil
	}

	return
}

func (lur *LicenseUploadReq) filename(filename string, name string, code string) (final string, err error) {
	if name, ee := envsubst.Eval(filename, lur.env(name, code)); nil != ee {
		err = ee
	} else if LicenseTypeDirect == lur.Type {
		final = filepath.Join(lur.Input, name)
	} else {
		final = filepath.Join(lur.Output, name)
	}

	if LicenseTypeDirect == lur.Type {
		if _, exists := gfx.Exists(final); !exists {
			final = ""
			err = exc.NewFields("文件不存在", field.New("企业名称", name), field.New("统一代码", code))
		}
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
