package pagination

import (
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
)

type HateoasPaginationStrategy struct {
	NextLinkPath string
	DataStrategy strategies.DataStrategy
}

func (s HateoasPaginationStrategy) ApplyPagination(req http.Request, previousResponse *[]byte) http.Request {
	if previousResponse != nil {
		newReq, _ := http.NewRequest(req.Method, *s.DataStrategy.GetSingleValue(*previousResponse, s.NextLinkPath), req.Body)
		req = *newReq
	}
	return req

}

func (s HateoasPaginationStrategy) HasMoreData(response []byte) bool {
	nextLink := s.DataStrategy.GetSingleValue(response, s.NextLinkPath)
	if nextLink == nil {
		return false
	} else {
		return true
	}
}
