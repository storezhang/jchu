package args

type TokenServer struct {
	Addr  string `validate:"required"`
	Token string `validate:"required"`
}

func newTokenServer() *TokenServer {
	return new(TokenServer)
}
