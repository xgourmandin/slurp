package slurp

import (
	"github.com/xgourmandin/slurp/internal/core/usecases"
	"github.com/xgourmandin/slurp/internal/handlers"
)

func NewSlurpEngine() usecases.SlurpAnApiUseCase {
	return usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}
}
