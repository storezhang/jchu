package ft

type licenseArgs struct {
	Enterprise string `validate:"required"`
	Input      string `validate:"required"`
	Type       string `validate:"oneof=word direct pdf"`
	Output     string `validate:"required"`
	Filename   string `validate:"required"`
	Skipped    int    `validate:"required"`
	Sheet      string `validate:"required"`
}

func newLicenseArgs() *licenseArgs {
	return &licenseArgs{
		Enterprise: `enterprise.xlsx`,
		Type:       "word",
		Output:     `license`,
		Filename:   "${NAME}-${CODE}.docx",
		Skipped:    1,
		Sheet:      `Sheet1`,
	}
}
