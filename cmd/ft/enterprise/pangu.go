package enterprise

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newCommand,

		newTemplate,
	)
}
