package domain

import (
	"net/http"
	"slurp/internal/core/domain/strategies"
)

type Context struct {
	ApiConfig              ApiConfiguration
	HttpStrategy           strategies.HttpStrategy
	PaginationStrategy     strategies.PaginationStrategy
	AuthenticationStrategy strategies.AuthenticationStrategy
	DataStrategy           strategies.DataStrategy
}

func (c Context) CreateRequest() (*http.Request, error) {
	request, err := c.HttpStrategy.CreateRequest(c.ApiConfig.Url)
	if err != nil {
		return nil, err
	}
	paginated := c.PaginationStrategy.ApplyPagination(*request)
	authenticated := c.AuthenticationStrategy.AddAuthentication(paginated)
	return &authenticated, nil
}
