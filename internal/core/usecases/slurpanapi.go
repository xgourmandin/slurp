package usecases

import (
	"fmt"
	"slurp/internal/core/domain"
	"slurp/internal/core/ports"
	"slurp/internal/handlers/strategies"
)

type SlurpAnApiUseCase struct {
	ReqHandler ports.RequestHandler
}

func (s SlurpAnApiUseCase) SlurpAPI(ctx domain.Context) {
	ctx.HttpStrategy = strategies.CreateHttpStrategy(ctx.ApiConfig.Method)
	dataStrategy := strategies.CreateDataStrategy(ctx.ApiConfig.DataType, ctx.ApiConfig.DataRoot)
	ctx.PaginationStrategy = strategies.CreatePaginationStrategy(ctx.ApiConfig, dataStrategy)
	ctx.DataStrategy = dataStrategy

	hasMore := true
	for hasMore {
		response := s.ReqHandler.SendRequest(ctx)
		out := make(chan interface{})
		go ctx.DataStrategy.ExtractData(response, out)
		for v := range out {
			fmt.Println(v)
		}
		hasMore = ctx.PaginationStrategy.HasMoreData(response)
	}

}
