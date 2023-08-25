package authentication

import "net/http"

type NoAuthenticationStrategy struct {
}

func (NoAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	return req
}
