package pagination

import (
	"net/http"
	"slurp/internal/core/ports/strategies"
	"strconv"
)

type OffsetLimitPaginationStrategy struct {
	OffsetParam   string
	LimitParam    string
	CurrentOffset int
	LimitValue    int
	DataStrategy  strategies.DataStrategy
}

func (s *OffsetLimitPaginationStrategy) ApplyPagination(req http.Request, _ *[]byte) http.Request {
	q := req.URL.Query()
	q.Set(s.OffsetParam, strconv.Itoa(s.CurrentOffset))
	q.Set(s.LimitParam, strconv.Itoa(s.LimitValue))
	req.URL.RawQuery = q.Encode()
	s.CurrentOffset = s.CurrentOffset + s.LimitValue
	return req
}

func (s *OffsetLimitPaginationStrategy) HasMoreData(response []byte) bool {
	return s.DataStrategy.GetResultSize(response) == s.LimitValue
}
