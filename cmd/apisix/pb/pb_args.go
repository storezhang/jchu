package pb

type pbArgs struct {
	Id string `validate:"required"`
}

func newPbArgs() *pbArgs {
	return new(pbArgs)
}
