package excel

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newMerge,
		newCommand,
	)
}
