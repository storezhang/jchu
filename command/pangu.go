package command

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newDuplicate,
	)
}
