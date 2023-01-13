package pb

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newCommand,

		newPbArgs,
		newUpload,
	)
}
