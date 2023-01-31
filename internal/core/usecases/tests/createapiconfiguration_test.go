package tests

import (
	"fmt"
	"os"
	"slurp/internal/core/usecases"
	"testing"
)

func TestApiConfigurationCreation(t *testing.T) {
	configurationUseCase := usecases.CreateApiConfigurationUseCase{
		ApiRepository: MockApiRepository{},
	}

	configuration, _ := configurationUseCase.CreateApiConfiguration("test_api")

	if configuration.Url != "https://test/api.com" {
		t.Errorf("Wrong URL parsed from YAML")
	}
	if configuration.Method != "GET" {
		t.Errorf("Wrong Http method parsed from YAML")
	}
}

func TestWrongMethodError(t *testing.T) {
	configurationUseCase := usecases.CreateApiConfigurationUseCase{
		ApiRepository: MockApiRepository{},
	}

	configuration, err := configurationUseCase.CreateApiConfiguration("wrong_method")

	if configuration != nil {
		t.Errorf("Wrong configuration shall not initialize a struct")
	}
	if err.Error() != "Wrong HTTP method given: OPTION" {
		t.Errorf("a correct error shall be returned in case of wrong http method configuration")
	}
}

func TestPaginatedApi(t *testing.T) {
	configurationUseCase := usecases.CreateApiConfigurationUseCase{
		ApiRepository: MockApiRepository{},
	}

	configuration, err := configurationUseCase.CreateApiConfiguration("paginated_api")

	if err != nil {
		t.Errorf("Configuration shall be correctly created")
	}
	if configuration.PaginationConfig.PaginationType != "PAGE_LIMIT" {
		t.Errorf("Wrong pagination type")
	}
	if configuration.PaginationConfig.PageParam != "page" {
		t.Errorf("Wrong pagination page param")
	}
	if configuration.PaginationConfig.LimitParam != "per_page" {
		t.Errorf("Wrong pagination limit param")
	}
	if configuration.PaginationConfig.PageSize != 25 {
		t.Errorf("Wrong pagination limit size")
	}
}

type MockApiRepository struct {
}

func (MockApiRepository) GetApiConfiguration(name string) ([]byte, error) {
	file, err := os.ReadFile(fmt.Sprintf("./%s.yaml", name))
	return file, err

}
