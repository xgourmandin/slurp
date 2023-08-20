package secrets

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"context"
)

type GcpSecretManager struct {
}

func (m GcpSecretManager) GetSecretValue(secretName string) (string, error) {
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		return "", err
	}
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", err
	} else {
		return string(result.Payload.Data), nil
	}
}

func getClient(ctx context.Context) (*secretmanager.Client, error) {
	return secretmanager.NewClient(ctx)
}
