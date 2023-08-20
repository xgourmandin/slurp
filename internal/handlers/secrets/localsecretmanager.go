package secrets

import (
	"fmt"
	"os"
)

type LocalSecretManager struct {
}

func (m LocalSecretManager) GetSecretValue(secretName string) (string, error) {
	content, err := os.ReadFile(fmt.Sprintf("./%s.secret", secretName))
	if err != nil {
		return "", err
	} else {
		return string(content), nil
	}
}
