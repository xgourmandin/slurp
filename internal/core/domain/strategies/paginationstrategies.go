package strategies

import "net/http"

type PaginationStrategy interface {
	ApplyPagination(http.Request) http.Request
	HasMoreData(response []byte) bool
}
