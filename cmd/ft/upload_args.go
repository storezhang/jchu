package ft

type uploadArgs struct {
	command *args        `validate:"required"`
	license *licenseArgs `validate:"required"`
}
