package ft

type licenseArgs struct {
	enterprise string
	output     string
	skipped    int
	sheet      string
}

func newLicenseArgs() *licenseArgs {
	return &licenseArgs{
		enterprise: `enterprise.xlsx`,
		output:     `license`,
		skipped:    1,
		sheet:      `Sheet1`,
	}
}
