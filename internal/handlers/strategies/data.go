package strategies

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"reflect"
	"slurp/internal/core/ports/strategies"
)

type JsonDataStrategy struct {
	DataRootPath string // JSON Path of the root of the data to fetch. Can point to an array (nominal case) or a single element
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

func (s JsonDataStrategy) ExtractData(body []byte, out chan interface{}) error {
	jsonData := interface{}(nil)
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return err
	}
	root, err := jsonpath.Get(s.DataRootPath, jsonData)
	if err != nil {
		return err
	}
	switch root.(type) {
	case []interface{}:
		outputArrayData(root.([]interface{}), out)
	case interface{}:
		outputSingleValue(root, out)
	default:
		return errors.New(fmt.Sprintf("Unkonw data type for the confiogured API: %s", reflect.TypeOf(root)))
	}

	return nil
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

func CreateDataStrategy(dataType string, dataRoot string) strategies.DataStrategy {
	return JsonDataStrategy{DataRootPath: dataRoot}
}
