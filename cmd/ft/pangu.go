package ft

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Dependencies(
		newArgs,
		newCommand,

		newUpload,
		newLicenseArgs,
		newLicense,
		newEnterprise,

		newTemplate,
	)
}
