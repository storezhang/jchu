package license

import (
	"github.com/storezhang/cli/args"
)

type uploadArgs struct {
	ft      *args.Ft      `validate:"required"`
	license *args.License `validate:"required"`
}
