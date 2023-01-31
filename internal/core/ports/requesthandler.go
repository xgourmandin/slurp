package ports

import "slurp/internal/core/domain"

type RequestHandler interface {
	SendRequest(ctx domain.Context) []byte
}
