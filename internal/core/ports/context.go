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
	request = addAdditionalHeaders(request, c.ApiConfig.AdditionalHeaders)
	request = addAdditionalQueryParams(request, c.ApiConfig.AdditionalQueryParams)
	if err != nil {
		return nil, err
	}
	paginated := c.PaginationStrategy.ApplyPagination(*request, c.PreviousResponse)
	authenticated := c.AuthenticationStrategy.AddAuthentication(paginated)
	return &authenticated, nil
}

func addAdditionalHeaders(request *http.Request, headers map[string]string) *http.Request {
	for k, v := range headers {
		request.Header.Add(k, v)
	}
	return request
}

func addAdditionalQueryParams(request *http.Request, params map[string]string) *http.Request {
	for k, v := range params {
		q := request.URL.Query()
		q.Set(k, v)
		request.URL.RawQuery = q.Encode()
	}
	return request
}
