package main

import (
	"log"
	"os"
	"slurp/internal/core/usecases"
	"slurp/internal/handlers"
	"slurp/internal/handlers/repositories"
)

func main() {

	ctx, err := usecases.CreateContextUseCase{
		ApiConfigurationRepository: repositories.LocalApiRepository{},
	}.CreateContext(os.Args[1])
	if err != nil {
		log.Fatalf("An error ahs occured during API configuration parsing: %v", err)
		panic(1)
	}
	usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}.SlurpAPI(*ctx)

}
