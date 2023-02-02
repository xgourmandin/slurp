package usecases

import (
	"fmt"
	"slurp/internal/core/domain"
	"slurp/internal/core/ports"
	"slurp/internal/handlers"
	"slurp/internal/handlers/strategies"
)

type CreateContextUseCase struct {
}

func (c CreateContextUseCase) CreateContext(apiConfig domain.ApiConfiguration) ports.Context {
	ctx := ports.Context{
		ApiConfig: apiConfig,
	}
	ctx.HttpStrategy = strategies.CreateHttpStrategy(ctx.ApiConfig.Method)
	dataStrategy := strategies.CreateDataStrategy(ctx.ApiConfig.DataType, ctx.ApiConfig.DataRoot)
	ctx.PaginationStrategy = strategies.CreatePaginationStrategy(ctx.ApiConfig, dataStrategy)
	ctx.AuthenticationStrategy = strategies.CreateAuthenticationStrategy(ctx.ApiConfig)
	ctx.DataStrategy = dataStrategy

	switch apiConfig.OutputConfig.OutputType {
	case "FILE":
		ctx.ApiDataWriter = handlers.FileWriter{
			Format:   "json",
			FileName: apiConfig.OutputConfig.FileName,
		}
	case "BUCKET":
		ctx.ApiDataWriter = handlers.GcsStorageWriter{
			Format:     "json",
			BucketName: apiConfig.OutputConfig.BucketName,
			FileName:   apiConfig.OutputConfig.FileName,
		}
	case "BIGQUERY":
		ctx.ApiDataWriter = handlers.NewBigQueryWriter(
			apiConfig.OutputConfig.Project,
			apiConfig.OutputConfig.Dataset,
			apiConfig.OutputConfig.Table,
			apiConfig.OutputConfig.Autodetect,
			fmt.Sprintf("/tmp/slurp-%s.json", apiConfig.OutputConfig.Table),
		)

	default:
		ctx.ApiDataWriter = handlers.LogWriter{}
	}
	return ctx
}
