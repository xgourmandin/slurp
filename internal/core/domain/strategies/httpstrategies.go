package strategies

import (
	"net/http"
)

type HttpStrategy interface {
	CreateRequest(url string) (*http.Request, error)
}
