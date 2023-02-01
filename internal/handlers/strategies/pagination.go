package strategies

import (
	"net/http"
	"slurp/internal/core/domain"
	"slurp/internal/core/domain/strategies"
	"strconv"
)

type NoPaginationStrategy struct {
}

func (s NoPaginationStrategy) ApplyPagination(req http.Request) http.Request {
	return req
}

func (NoPaginationStrategy) HasMoreData(_ []byte) bool {
	return false
}

type PageLimitPaginationStrategy struct {
	PageParam     string
	LimitParam    string
	CurrentPage   int
	LimitValue    int
	MoreItemsPath *string // Data selector Path of the field telling us if there is more data to be fetched in the future (can be null, in this case, we rely on the size of the API response compared to the limit configured)
	dataStrategy  strategies.DataStrategy
}

func (s *PageLimitPaginationStrategy) ApplyPagination(req http.Request) http.Request {
	q := req.URL.Query()
	q.Set(s.PageParam, strconv.Itoa(s.CurrentPage))
	q.Set(s.LimitParam, strconv.Itoa(s.LimitValue))
	req.URL.RawQuery = q.Encode()
	s.CurrentPage = s.CurrentPage + 1
	return req
}

func (s *PageLimitPaginationStrategy) HasMoreData(response []byte) bool {
	if s.MoreItemsPath == nil {
		return s.dataStrategy.GetResultSize(response) == s.LimitValue
	} else {
		// TODO: Implement the more item field behaviour
		return false
	}
}

type OffsetLimitPaginationStrategy struct {
	OffsetParam   string
	LimitParam    string
	CurrentOffset int
	LimitValue    int
	dataStrategy  strategies.DataStrategy
}

func (s *OffsetLimitPaginationStrategy) ApplyPagination(req http.Request) http.Request {
	q := req.URL.Query()
	q.Set(s.OffsetParam, strconv.Itoa(s.CurrentOffset))
	q.Set(s.LimitParam, strconv.Itoa(s.LimitValue))
	req.URL.RawQuery = q.Encode()
	s.CurrentOffset = s.CurrentOffset + s.LimitValue
	return req
}

func (s *OffsetLimitPaginationStrategy) HasMoreData(response []byte) bool {
	return s.dataStrategy.GetResultSize(response) == s.LimitValue
}

func CreatePaginationStrategy(apiConfig domain.ApiConfiguration, dataStrategy strategies.DataStrategy) strategies.PaginationStrategy {
	switch apiConfig.PaginationConfig.PaginationType {
	case "PAGE_LIMIT":
		return &PageLimitPaginationStrategy{
			PageParam:     apiConfig.PaginationConfig.PageParam,
			LimitParam:    apiConfig.PaginationConfig.LimitParam,
			CurrentPage:   1,
			LimitValue:    apiConfig.PaginationConfig.PageSize,
			MoreItemsPath: nil,
			dataStrategy:  dataStrategy,
		}
	case "OFFSET_LIMIT":
		return &OffsetLimitPaginationStrategy{
			OffsetParam:   apiConfig.PaginationConfig.PageParam,
			LimitParam:    apiConfig.PaginationConfig.LimitParam,
			CurrentOffset: 0,
			LimitValue:    apiConfig.PaginationConfig.PageSize,
			dataStrategy:  dataStrategy,
		}
	default:
		return NoPaginationStrategy{}
	}
}
