package main

import (
	"github.com/xgourmandin/slurp/internal/core/usecases"
	"github.com/xgourmandin/slurp/internal/handlers"
	"github.com/xgourmandin/slurp/internal/handlers/repositories"
	"log"
	"os"
)

func main() {

	ctx, err := usecases.CreateContextUseCase{
		ApiConfigurationRepository: repositories.LocalApiRepository{},
	}.CreateContext(os.Args[1])
	if err != nil {
		log.Fatalf("An error ahs occured during API configuration parsing: %v", err)
	}
	usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}.SlurpAPI(*ctx)

}
