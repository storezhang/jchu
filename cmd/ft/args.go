package ft

type args struct {
	id     string `validate:"required"`
	key    string `validate:"required"`
	secret string `validate:"required"`
	addr   string `validate:"required"`
	result string `validate:"required"`
}

func newArgs() *args {
	return &args{
		addr:   "https://202.61.91.57:8092",
		result: "result.txt",
	}
}
