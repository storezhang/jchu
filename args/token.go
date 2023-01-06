package args

type Token struct {
	Addr  string `validate:"required"`
	Token string `validate:"required"`
}

func newToken() *Token {
	return new(Token)
}
