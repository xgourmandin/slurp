package strategies

import (
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"github.com/xgourmandin/slurp/internal/handlers/strategies/authentication"
)

func CreateAuthenticationStrategy(apiConfig configuration.ApiConfiguration, manager ports.SecretManager) strategies.AuthenticationStrategy {
	switch apiConfig.AuthConfig.AuthType {
	case "API_KEY":
		return authentication.ApiTokenAuthenticationStrategy{
			SecretManager: manager,
			Token:         apiConfig.AuthConfig.TokenSecret,
			InHeader:      apiConfig.AuthConfig.InHeader,
			AuthParam:     apiConfig.AuthConfig.TokenParam,
		}
	case "CLIENT_CREDS":
		return &authentication.ClientCredentialsAuthenticationStrategy{
			SecretManager:   manager,
			AccessTokenUrl:  apiConfig.AuthConfig.AccessTokenUrl,
			PayloadTemplate: apiConfig.AuthConfig.PayloadTemplate,
			ClientId:        apiConfig.AuthConfig.ClientIdSecret,
			ClientSecret:    apiConfig.AuthConfig.ClientSecretSecret,
			AccessTokenPath: apiConfig.AuthConfig.AccessTokenPath,
		}
	default:
		return authentication.NoAuthenticationStrategy{}
	}
}
