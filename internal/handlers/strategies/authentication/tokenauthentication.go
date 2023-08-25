package authentication

import (
	"fmt"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"net/http"
)

type ApiTokenAuthenticationStrategy struct {
	SecretManager ports.SecretManager
	Token         string // Not the token itself, but the env variable name that holds the token
	InHeader      bool   // If true, the token goes in an Authentication header, else, it goes in the query string
	AuthParam     string
}

func (s ApiTokenAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	tokenValue, err := s.SecretManager.GetSecretValue(s.Token)
	if err != nil {
		fmt.Printf("%v\n", err)
		return req
	}
	if s.InHeader {
		req.Header.Add(s.AuthParam, tokenValue)
	} else {
		q := req.URL.Query()
		q.Set(s.AuthParam, tokenValue)
		req.URL.RawQuery = q.Encode()
	}
	return req
}
