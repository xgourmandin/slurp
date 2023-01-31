package repositories

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

type GcpStorageApiConfigurationRepository struct {
	apiConfigurationBucket string
}

func (r GcpStorageApiConfigurationRepository) initStorageClient(ctx context.Context) (*storage.Client, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}

func (r GcpStorageApiConfigurationRepository) GetApiConfiguration(apiName string) ([]byte, error) {
	ctx := context.Background()
	client, err := r.initStorageClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	reader, err := client.Bucket(r.apiConfigurationBucket).Object(fmt.Sprintf("%s.yaml", apiName)).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%q).NewReader: %v", apiName, err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}
	return data, nil
}

type LocalApiRepository struct {
}

func (LocalApiRepository) GetApiConfiguration(name string) ([]byte, error) {
	file, err := os.ReadFile(fmt.Sprintf("./%s.yaml", name))
	return file, err

}
