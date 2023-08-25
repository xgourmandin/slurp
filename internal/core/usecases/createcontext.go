package usecases

import (
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"github.com/xgourmandin/slurp/internal/handlers/secrets"
	"github.com/xgourmandin/slurp/internal/handlers/strategies"
)

type CreateContextUseCase struct {
	ApiConfigurationRepository ports.ApiConfigurationRepository
}

func (c CreateContextUseCase) CreateContextFromConfig(configuration *configuration.ApiConfiguration, env string) (*ports.Context, error) {
	ctx := ports.Context{
		ApiConfig: *configuration,
	}
	ctx.HttpStrategy = strategies.CreateHttpStrategy(ctx.ApiConfig.Method)
	dataStrategy := strategies.CreateDataStrategy(ctx.ApiConfig.DataConfig)
	ctx.PaginationStrategy = strategies.CreatePaginationStrategy(ctx.ApiConfig, dataStrategy)
	secretManager := secrets.NewSecretManager(env)
	ctx.AuthenticationStrategy = strategies.CreateAuthenticationStrategy(ctx.ApiConfig, secretManager)
	ctx.ApiDataWriter = strategies.NewWriterStrategy(ctx.ApiConfig)
	ctx.DataStrategy = dataStrategy
	return &ctx, nil
}

func (c CreateContextUseCase) CreateContext(apiName string) (*ports.Context, error) {
	config, err := c.ApiConfigurationRepository.GetApiConfiguration(apiName)
	if err != nil {
		return nil, err
	}
	return c.CreateContextFromConfig(config, "SERVER")
}
