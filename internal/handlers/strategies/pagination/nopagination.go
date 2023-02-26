package pagination

import "net/http"

type NoPaginationStrategy struct {
}

func (s NoPaginationStrategy) ApplyPagination(req http.Request, i *[]byte) http.Request {
	return req
}

func (NoPaginationStrategy) HasMoreData(_ []byte) bool {
	return false
}
