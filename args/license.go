package args

type License struct {
	Enterprise string   `validate:"required"`
	Input      string   `validate:"required"`
	Type       string   `validate:"oneof=word direct pdf"`
	Output     string   `validate:"required"`
	Filenames  []string `validate:"required"`
	Skipped    int      `validate:"required"`
	Sheet      string   `validate:"required"`
}

func newLicense() *License {
	return &License{
		Enterprise: `enterprise.xlsx`,
		Type:       "word",
		Output:     `license`,
		Filenames: []string{
			"${NAME}-${CODE}.pdf",
			"${NAME}+${CODE}.pdf",
			"${NAME}_${CODE}.pdf",
			"${NAME} ${CODE}.pdf",
			"${NAME}${CODE}.pdf",
			"${NAME}（${CODE}）.pdf",
			"${CODE}.pdf",
		},
		Skipped: 1,
		Sheet:   `Sheet1`,
	}
}
