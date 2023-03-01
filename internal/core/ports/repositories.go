package ports

import "github.com/xgourmandin/slurp"

type ApiConfigurationRepository interface {
	GetApiConfiguration(apiname string) (*slurp.ApiConfiguration, error)
}
