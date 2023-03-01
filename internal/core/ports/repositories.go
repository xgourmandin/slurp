package ports

import (
	"github.com/xgourmandin/slurp/configuration"
)

type ApiConfigurationRepository interface {
	GetApiConfiguration(apiname string) (*configuration.ApiConfiguration, error)
}
