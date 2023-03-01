package repositories

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/xgourmandin/slurp"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

type GcpStorageApiConfigurationRepository struct {
	ApiConfigurationBucket string
}

func (r GcpStorageApiConfigurationRepository) initStorageClient(ctx context.Context) (*storage.Client, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	return client, nil
}

func (r GcpStorageApiConfigurationRepository) GetApiConfiguration(apiName string) (*slurp.ApiConfiguration, error) {
	ctx := context.Background()
	client, err := r.initStorageClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	reader, err := client.Bucket(r.ApiConfigurationBucket).Object(fmt.Sprintf("%s.yaml", apiName)).NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("Object(%s.yaml).NewReader: %v", apiName, err)
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %v", err)
	}
	configuration := slurp.ApiConfiguration{}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return nil, err
	} else {
		return &configuration, nil
	}
}

type LocalApiRepository struct {
}

func (r LocalApiRepository) GetApiConfiguration(name string) (*slurp.ApiConfiguration, error) {
	content, err := os.ReadFile(fmt.Sprintf("./%s.yaml", name))
	if err != nil {
		return nil, err
	}
	configuration := slurp.ApiConfiguration{}
	err = yaml.Unmarshal(content, &configuration)
	if err != nil {
		return nil, err
	} else {
		return &configuration, nil
	}
}
