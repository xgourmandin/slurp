package strategies

import (
	"net/http"
	"os"
	"slurp/internal/core/ports"
	"slurp/internal/core/ports/strategies"
)

type NoAuthenticationStrategy struct {
}

func (NoAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	return req
}

type ApiTokenAuthenticationStrategy struct {
	Token     string // Not the token itself, but the env variable name that holds the token
	InHeader  bool   // If true, the token goes in an Authentication header, else, it goes in the query string
	AuthParam string
}

func (s ApiTokenAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	if s.InHeader {
		req.Header.Add(s.AuthParam, os.Getenv(s.Token))
	} else {
		q := req.URL.Query()
		q.Set(s.AuthParam, os.Getenv(s.Token))
		req.URL.RawQuery = q.Encode()
	}
	return req
}

func CreateAuthenticationStrategy(apiConfig ports.ApiConfiguration) strategies.AuthenticationStrategy {
	switch apiConfig.AuthConfig.AuthType {
	case "API_KEY":
		return ApiTokenAuthenticationStrategy{
			Token:     apiConfig.AuthConfig.TokenEnv,
			InHeader:  apiConfig.AuthConfig.InHeader,
			AuthParam: apiConfig.AuthConfig.TokenParam,
		}
	default:
		return NoAuthenticationStrategy{}
	}
}
