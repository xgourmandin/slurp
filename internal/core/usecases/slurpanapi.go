package usecases

import (
	"github.com/xgourmandin/slurp/internal/core/ports"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"log"
)

type SlurpAnApiUseCase struct {
	ReqHandler ports.RequestHandler
}

func (s SlurpAnApiUseCase) SlurpAPIWithState(ctx ports.Context, state strategies.PaginationState) (int, strategies.PaginationState) {
	ctx.PaginationStrategy = ctx.PaginationStrategy.Configure(state)
	dataCount := s.SlurpAPI(ctx)
	return dataCount, ctx.PaginationStrategy.NextContext()
}

func (s SlurpAnApiUseCase) SlurpAPI(ctx ports.Context) int {
	hasMore := true
	batchSizeReached := false
	dataCount := 0
	for hasMore && !batchSizeReached {
		response, err := s.ReqHandler.SendRequest(ctx)
		if err != nil {
			log.Printf("%v", err)
			return dataCount
		}
		ctx.PreviousResponse = &response
		out := make(chan interface{})
		go ctx.DataStrategy.ExtractData(response, out)
		for v := range out {
			ctx.ApiDataWriter = ctx.ApiDataWriter.StoreApiResult(v)
			dataCount++
		}
		hasMore = ctx.PaginationStrategy.HasMoreData(response)
		batchSizeReached = ctx.PaginationStrategy.IsBatchSizeReached()
	}
	err := ctx.ApiDataWriter.Finalize()
	if err != nil {
		log.Fatalf("An error has occured during output finalization: %v", err)
	}
	return dataCount

}
