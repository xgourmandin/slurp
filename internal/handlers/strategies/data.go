package strategies

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"reflect"
)

type JsonDataStrategy struct {
	DataRootPath string // JSON Path of the root of the data to fetch. Can point to an array (nominal case) or a single element
}

func (s JsonDataStrategy) GetSingleValue(body []byte, path string) *string {
	jsonData := interface{}(nil)
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil
	}
	root, err := jsonpath.Get(path, jsonData)
	if err != nil {
		return nil
	}
	if root == nil {
		return nil
	} else {
		value := root.(string)
		return &value
	}
}

func (s JsonDataStrategy) GetResultSize(response []byte) int {
	jsonData := interface{}(nil)
	err := json.Unmarshal(response, &jsonData)
	if err != nil {
		return 0
	}
	root, err := jsonpath.Get(s.DataRootPath, jsonData)
	if err != nil {
		return 0
	}
	return len(root.([]interface{}))
}

func (s JsonDataStrategy) ExtractData(body []byte, out chan interface{}) {
	fmt.Print(string(body))
	jsonData := interface{}(nil)
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		println(fmt.Sprintf("Error during JSON unmarshall : %s", err.Error()))
		close(out)
	}
	root, err := jsonpath.Get(s.DataRootPath, jsonData)
	if err != nil {
		println(fmt.Sprintf("Error during data extraction with using json path %s: %s", s.DataRootPath, err.Error()))
		close(out)
	}
	switch root.(type) {
	case []interface{}:
		outputArrayData(root.([]interface{}), out)
	case interface{}:
		outputSingleValue(root, out)
	default:
		println(fmt.Sprintf("Unkonw data type for the configured API: %s", reflect.TypeOf(root)))
	}
}

func outputArrayData(data []interface{}, out chan interface{}) {
	defer close(out)
	for _, v := range data {
		out <- v
	}
}

func outputSingleValue(data interface{}, out chan interface{}) {
	defer close(out)
	out <- data
}

func CreateDataStrategy(dataConfig configuration.DataConfiguration) strategies.DataStrategy {
	return JsonDataStrategy{DataRootPath: dataConfig.DataRoot}
}
