package strategies

import (
	"net/http"
)

type PaginationStrategy interface {
	ApplyPagination(http.Request, *[]byte) http.Request
	HasMoreData(response []byte) bool
	Configure(state PaginationState) PaginationStrategy
	IsBatchSizeReached() bool
	NextContext() PaginationState
}
