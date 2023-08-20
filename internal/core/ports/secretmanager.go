package ports

type SecretManager interface {
	GetSecretValue(secretName string) (string, error)
}
