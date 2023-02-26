package ports

import (
	"net/http"
	"slurp/internal/core/domain"
	"slurp/internal/core/ports/strategies"
)

type Context struct {
	ApiConfig              domain.ApiConfiguration
	HttpStrategy           strategies.HttpStrategy
	PaginationStrategy     strategies.PaginationStrategy
	AuthenticationStrategy strategies.AuthenticationStrategy
	DataStrategy           strategies.DataStrategy
	ApiDataWriter          ApiDataWriter
	PreviousResponse       *[]byte
}

func (c Context) CreateRequest() (*http.Request, error) {
	request, err := c.HttpStrategy.CreateRequest(c.ApiConfig.Url)
	if err != nil {
		return nil, err
	}
	paginated := c.PaginationStrategy.ApplyPagination(*request, c.PreviousResponse)
	authenticated := c.AuthenticationStrategy.AddAuthentication(paginated)
	return &authenticated, nil
}
