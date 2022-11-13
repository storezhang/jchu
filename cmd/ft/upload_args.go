package ft

import (
	"github.com/goexl/xiren"
)

type uploadArgs struct {
	command *args
	license *licenseArgs
}

func (ua *uploadArgs) validate() (err error) {
	if err = xiren.Struct(ua.command); nil != err {
		return
	}
	if err = xiren.Struct(ua.license); nil != err {
		return
	}

	return
}
