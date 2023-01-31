package tests

import (
	"slurp/internal/core/domain"
	"testing"
)

func TestConfigurationCreationFromMap(t *testing.T) {
	testMap := make(map[string]interface{})
	testMap["url"] = "https://test.api.com"
	testMap["method"] = "GET"
	dataMap := make(map[string]interface{})
	dataMap["type"] = "JSON"
	testMap["data"] = dataMap

	configuration := domain.ApiConfiguration{}
	configuration.FromMap(testMap)

	if configuration.Url != "https://test.api.com" {
		t.Errorf("URL shall be set in the API configuration")
	}
	if configuration.Method != "GET" {
		t.Errorf("Http Method shall be set in the API configuration")
	}
	if configuration.DataType != "JSON" {
		t.Errorf("Data type shall be set in the API configuration")
	}

}
