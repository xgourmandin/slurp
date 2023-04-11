package usecases

import (
	"fmt"
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"github.com/xgourmandin/slurp/internal/handlers"
	"github.com/xgourmandin/slurp/internal/handlers/strategies"
	"strings"
	"time"
)

type CreateContextUseCase struct {
	ApiConfigurationRepository ports.ApiConfigurationRepository
}

func (c CreateContextUseCase) CreateContextFromConfig(configuration *configuration.ApiConfiguration) (*ports.Context, error) {
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
		chunked := strings.Split(configuration.OutputConfig.FileName, ".")
		filename := strings.Join(chunked[:len(chunked)-1], ".") + "-" + time.Now().Format("20060201150405") + "." + chunked[len(chunked)-1]
		ctx.ApiDataWriter = handlers.GcsStorageWriter{
			Format:     "json",
			BucketName: configuration.OutputConfig.BucketName,
			FileName:   filename,
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

func (c CreateContextUseCase) CreateContext(apiName string) (*ports.Context, error) {
	config, err := c.ApiConfigurationRepository.GetApiConfiguration(apiName)
	if err != nil {
		return nil, err
	}
	return c.CreateContextFromConfig(config)
}
