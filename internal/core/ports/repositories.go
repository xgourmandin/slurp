package ports

type ApiConfigurationRepository interface {
	GetApiConfiguration(apiname string) (*ApiConfiguration, error)
}
