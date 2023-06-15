package middleware

type Middleware struct{}

func New() *Middleware {
	return new(Middleware)
}
