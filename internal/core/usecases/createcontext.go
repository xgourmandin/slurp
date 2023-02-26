package usecases

import (
	"fmt"
	"slurp/internal/core/ports"
	"slurp/internal/handlers"
	"slurp/internal/handlers/strategies"
)

type CreateContextUseCase struct {
	ApiConfigurationRepository ports.ApiConfigurationRepository
}

func (c CreateContextUseCase) CreateContext(apiName string) (*ports.Context, error) {
	configuration, err := c.ApiConfigurationRepository.GetApiConfiguration(apiName)
	if err != nil {
		return nil, err
	}
	ctx := ports.Context{
		ApiConfig: *configuration,
	}
	ctx.HttpStrategy = strategies.CreateHttpStrategy(ctx.ApiConfig.Method)
	dataStrategy := strategies.CreateDataStrategy(ctx.ApiConfig.DataConfig)
	ctx.PaginationStrategy = strategies.CreatePaginationStrategy(ctx.ApiConfig, dataStrategy)
	ctx.AuthenticationStrategy = strategies.CreateAuthenticationStrategy(ctx.ApiConfig)
	ctx.DataStrategy = dataStrategy

	switch configuration.OutputConfig.OutputType {
	case "FILE":
		ctx.ApiDataWriter = handlers.FileWriter{
			Format:   "json",
			FileName: configuration.OutputConfig.FileName,
		}
	case "BUCKET":
		ctx.ApiDataWriter = handlers.GcsStorageWriter{
			Format:     "json",
			BucketName: configuration.OutputConfig.BucketName,
			FileName:   configuration.OutputConfig.FileName,
		}
	case "BIGQUERY":
		ctx.ApiDataWriter = handlers.NewBigQueryWriter(
			configuration.OutputConfig.Project,
			configuration.OutputConfig.Dataset,
			configuration.OutputConfig.Table,
			configuration.OutputConfig.Autodetect,
			fmt.Sprintf("/tmp/slurp-%s.json", configuration.OutputConfig.Table),
		)

	default:
		ctx.ApiDataWriter = handlers.LogWriter{}
	}
	return &ctx, nil
}
