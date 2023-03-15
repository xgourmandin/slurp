package ports

type RequestHandler interface {
	SendRequest(ctx Context) ([]byte, error)
}
