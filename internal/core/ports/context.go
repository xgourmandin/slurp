package ports

import (
	"github.com/xgourmandin/slurp"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
)

type Context struct {
	ApiConfig              slurp.ApiConfiguration
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
