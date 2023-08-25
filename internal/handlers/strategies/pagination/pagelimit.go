package pagination

import (
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
	"strconv"
)

type PageLimitPaginationStrategy struct {
	PageParam     string
	LimitParam    string
	CurrentPage   int
	LimitValue    int
	MoreItemsPath *string // Data selector Path of the field telling us if there is more data to be fetched in the future (can be null, in this case, we rely on the size of the API response compared to the limit configured)
	DataStrategy  strategies.DataStrategy
	batchSize     int
	pageCount     int
}

func (s *PageLimitPaginationStrategy) ApplyPagination(req http.Request, _ *[]byte) http.Request {
	q := req.URL.Query()
	q.Set(s.PageParam, strconv.Itoa(s.CurrentPage))
	q.Set(s.LimitParam, strconv.Itoa(s.LimitValue))
	req.URL.RawQuery = q.Encode()
	s.CurrentPage = s.CurrentPage + 1
	s.pageCount = s.pageCount + 1
	return req
}

func (s *PageLimitPaginationStrategy) HasMoreData(response []byte) bool {
	if s.MoreItemsPath == nil {
		return s.DataStrategy.GetResultSize(response) == s.LimitValue
	} else {
		// TODO: Implement the more item field behaviour
		return false
	}
}

func (s *PageLimitPaginationStrategy) Configure(ctx strategies.PaginationState) strategies.PaginationStrategy {
	s.CurrentPage = ctx.NextValue.(int)
	s.batchSize = ctx.BatchSize
	return s
}

func (s *PageLimitPaginationStrategy) IsBatchSizeReached() bool {
	return s.batchSize == s.pageCount
}

func (s PageLimitPaginationStrategy) NextContext() strategies.PaginationState {
	return strategies.PaginationState{
		BatchSize: s.batchSize,
		NextValue: s.CurrentPage + 1,
	}
}
