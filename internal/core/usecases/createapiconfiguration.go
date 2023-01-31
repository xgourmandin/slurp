package usecases

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"slurp/internal/core/domain"
	"slurp/internal/core/ports"
)

type CreateApiConfigurationUseCase struct {
	ApiRepository ports.ApiConfigurationRepository
}

func (uc CreateApiConfigurationUseCase) CreateApiConfiguration(apiName string) (*domain.ApiConfiguration, error) {
	configuration, err := uc.ApiRepository.GetApiConfiguration(apiName)
	if err != nil {
		return nil, err
	}
	var apiConfig map[string]interface{}
	if err := yaml.Unmarshal(configuration, &apiConfig); err != nil {
		fmt.Errorf("unable to parse YAML file: %v", err)
		return nil, err
	}
	apiConfiguration := domain.ApiConfiguration{}
	if err := apiConfiguration.FromMap(apiConfig); err != nil {
		return nil, err
	}
	return &apiConfiguration, nil
}
