package pagination

import (
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
	"strconv"
)

type OffsetLimitPaginationStrategy struct {
	OffsetParam   string
	LimitParam    string
	CurrentOffset int
	LimitValue    int
	DataStrategy  strategies.DataStrategy
	batchSize     int
	pageCount     int
}

func (s *OffsetLimitPaginationStrategy) ApplyPagination(req http.Request, _ *[]byte) http.Request {
	q := req.URL.Query()
	q.Set(s.OffsetParam, strconv.Itoa(s.CurrentOffset))
	q.Set(s.LimitParam, strconv.Itoa(s.LimitValue))
	req.URL.RawQuery = q.Encode()
	s.CurrentOffset = s.CurrentOffset + s.LimitValue
	s.pageCount = s.pageCount + 1
	return req
}

func (s *OffsetLimitPaginationStrategy) HasMoreData(response []byte) bool {
	return s.DataStrategy.GetResultSize(response) == s.LimitValue
}

func (s *OffsetLimitPaginationStrategy) Configure(ctx strategies.PaginationState) strategies.PaginationStrategy {
	s.CurrentOffset = ctx.NextValue.(int)
	s.batchSize = ctx.BatchSize
	return s
}

func (s *OffsetLimitPaginationStrategy) IsBatchSizeReached() bool {
	return s.batchSize == s.pageCount
}

func (s OffsetLimitPaginationStrategy) NextContext() strategies.PaginationState {
	return strategies.PaginationState{
		BatchSize: s.batchSize,
		NextValue: s.CurrentOffset + s.LimitValue,
	}
}
