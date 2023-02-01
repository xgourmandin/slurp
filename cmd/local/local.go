package main

import (
	"log"
	"os"
	"slurp/internal/core/usecases"
	"slurp/internal/handlers"
	"slurp/internal/repositories"
)

func main() {

	apiConfigUc := usecases.CreateApiConfigurationUseCase{
		ApiRepository: repositories.LocalApiRepository{},
	}

	apiConfiguration, err := apiConfigUc.CreateApiConfiguration(os.Args[1])
	if err != nil {
		log.Fatalf("An error ahs occured during API configuration parsing: %v", err)
		panic(1)
	}
	ctx := usecases.CreateContextUseCase{}.CreateContext(*apiConfiguration)

	usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}.SlurpAPI(ctx)

}
