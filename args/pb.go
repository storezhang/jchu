package args

type Pb struct {
	Id string `validate:"required"`
}

func newPb() *Pb {
	return new(Pb)
}
