package ports

type ApiConfigurationRepository interface {
	GetApiConfiguration(apiname string) ([]byte, error)
}
