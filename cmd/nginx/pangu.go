package nginx

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newCommand,
	)
}
