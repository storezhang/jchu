package ft

type args struct {
	Id     string `validate:"required"`
	Key    string `validate:"required"`
	Secret string `validate:"required"`
	Addr   string `validate:"required"`
	Result string `validate:"required"`
}

func newArgs() *args {
	return &args{
		Addr:   "https://202.61.91.57:8092",
		Result: "result.txt",
	}
}
