package secrets

import "github.com/xgourmandin/slurp/internal/core/ports"

func NewSecretManager(env string) ports.SecretManager {
	if env == "LOCAL" {
		return LocalSecretManager{}
	} else {
		return GcpSecretManager{}
	}
}
