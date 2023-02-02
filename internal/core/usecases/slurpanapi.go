package usecases

import (
	"log"
	"slurp/internal/core/ports"
)

type SlurpAnApiUseCase struct {
	ReqHandler ports.RequestHandler
}

func (s SlurpAnApiUseCase) SlurpAPI(ctx ports.Context) {

	hasMore := true
	for hasMore {
		response := s.ReqHandler.SendRequest(ctx)
		out := make(chan interface{})
		go ctx.DataStrategy.ExtractData(response, out)
		for v := range out {
			ctx.ApiDataWriter.StoreApiResult(v)
		}
		hasMore = ctx.PaginationStrategy.HasMoreData(response)
	}
	err := ctx.ApiDataWriter.Finalize()
	if err != nil {
		log.Fatalf("An error has occured during output finalization: %v", err)
	}

}
