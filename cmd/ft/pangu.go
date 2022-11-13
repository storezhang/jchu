package ft

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newCommand,

		newLicense,
		newUpload,
		newEnterprise,

		newTemplate,
	)
}
