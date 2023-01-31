package strategies

import (
	"net/http"
	"slurp/internal/core/domain/strategies"
	"strings"
)

type HttpGetStrategy struct {
}

func (s HttpGetStrategy) CreateRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type HttpPostStrategy struct {
}

func (s HttpPostStrategy) CreateRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func CreateHttpStrategy(method string) strategies.HttpStrategy {
	switch strings.ToLower(method) {
	case "get":
		return HttpGetStrategy{}
	case "post":
		return HttpPostStrategy{}
	default:
		return nil
	}
}
