package ports

import (
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
)

type Context struct {
	ApiConfig              configuration.ApiConfiguration
	HttpStrategy           strategies.HttpStrategy
	PaginationStrategy     strategies.PaginationStrategy
	AuthenticationStrategy strategies.AuthenticationStrategy
	DataStrategy           strategies.DataStrategy
	ApiDataWriter          strategies.WriterStrategy
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
