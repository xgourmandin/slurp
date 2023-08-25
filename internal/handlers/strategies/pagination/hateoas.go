package pagination

import (
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
)

type HateoasPaginationStrategy struct {
	NextLinkPath  string
	DataStrategy  strategies.DataStrategy
	nextLinkValue *string
	pageCount     int
	batchSize     int
}

func (s *HateoasPaginationStrategy) ApplyPagination(req http.Request, previousResponse *[]byte) http.Request {
	nextLink := ""
	if s.nextLinkValue != nil {
		nextLink = *s.nextLinkValue
	}
	if previousResponse != nil {
		nextLink = *s.DataStrategy.GetSingleValue(*previousResponse, s.NextLinkPath)
		s.nextLinkValue = &nextLink
		newReq, _ := http.NewRequest(req.Method, nextLink, req.Body)
		req = *newReq
	}
	s.pageCount = s.pageCount + 1
	return req

}

func (s *HateoasPaginationStrategy) HasMoreData(response []byte) bool {
	nextLink := s.DataStrategy.GetSingleValue(response, s.NextLinkPath)
	s.nextLinkValue = nextLink
	if nextLink == nil {
		return false
	} else {
		return true
	}
}

func (s *HateoasPaginationStrategy) Configure(ctx strategies.PaginationState) strategies.PaginationStrategy {
	nextLink := ctx.NextValue.(string)
	s.nextLinkValue = &nextLink
	s.batchSize = ctx.BatchSize
	return s
}

func (s *HateoasPaginationStrategy) IsBatchSizeReached() bool {
	return !(s.batchSize == 0) || s.batchSize == s.pageCount
}

func (s HateoasPaginationStrategy) NextContext() strategies.PaginationState {
	return strategies.PaginationState{
		BatchSize: s.batchSize,
		NextValue: s.nextLinkValue,
	}
}
