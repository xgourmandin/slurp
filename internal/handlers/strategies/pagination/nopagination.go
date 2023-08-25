package pagination

import (
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"net/http"
)

type NoPaginationStrategy struct {
}

func (s NoPaginationStrategy) ApplyPagination(req http.Request, i *[]byte) http.Request {
	return req
}

func (NoPaginationStrategy) HasMoreData(_ []byte) bool {
	return false
}

func (s NoPaginationStrategy) Configure(_ strategies.PaginationState) strategies.PaginationStrategy {
	return s
}

func (s NoPaginationStrategy) IsBatchSizeReached() bool {
	return true
}

func (s NoPaginationStrategy) NextContext() strategies.PaginationState {
	return strategies.PaginationState{}
}
