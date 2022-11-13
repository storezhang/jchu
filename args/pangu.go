package args

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newFt,
		newLicense,
	)
}
