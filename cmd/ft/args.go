package ft

type args struct {
	id     string
	key    string
	secret string
	addr   string
	result string
}

func newArgs() *args {
	return &args{
		addr:   "https://202.61.91.57:8092",
		result: "result.txt",
	}
}
