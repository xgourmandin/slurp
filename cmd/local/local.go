package main

import (
	"os"
	"slurp/internal/core/domain"
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
		panic(1)
	}
	ctx := domain.Context{
		ApiConfig: *apiConfiguration,
	}

	usecases.SlurpAnApiUseCase{ReqHandler: handlers.HttpHandler{}}.SlurpAPI(ctx)

}
