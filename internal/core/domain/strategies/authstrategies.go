package strategies

import "net/http"

type AuthenticationStrategy interface {
	AddAuthentication(http.Request) http.Request
}
