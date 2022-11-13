package args

type Ft struct {
	Id     string `validate:"required"`
	Key    string `validate:"required"`
	Secret string `validate:"required"`
	Addr   string `validate:"required"`
	Result string `validate:"required"`
}

func newFt() *Ft {
	return &Ft{
		Addr:   "https://202.61.91.57:8092",
		Result: "result.txt",
	}
}
