package usecases

import (
	"github.com/xgourmandin/slurp/internal/core/ports"
	"log"
)

type SlurpAnApiUseCase struct {
	ReqHandler ports.RequestHandler
}

func (s SlurpAnApiUseCase) SlurpAPI(ctx ports.Context) int {

	hasMore := true
	dataCount := 0
	for hasMore {
		response := s.ReqHandler.SendRequest(ctx)
		ctx.PreviousResponse = &response
		out := make(chan interface{})
		go ctx.DataStrategy.ExtractData(response, out)
		for v := range out {
			ctx.ApiDataWriter.StoreApiResult(v)
			dataCount++
		}
		hasMore = ctx.PaginationStrategy.HasMoreData(response)
	}
	err := ctx.ApiDataWriter.Finalize()
	if err != nil {
		log.Fatalf("An error has occured during output finalization: %v", err)
	}
	return dataCount

}
