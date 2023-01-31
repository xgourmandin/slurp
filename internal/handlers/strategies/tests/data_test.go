package tests

import (
	"os"
	"slurp/internal/handlers/strategies"
	"testing"
)

func TestJsonDataStrategy(t *testing.T) {
	strategy := strategies.JsonDataStrategy{DataRootPath: "$"}
	file, _ := os.ReadFile("./cats.json")
	out := make(chan interface{})
	go strategy.ExtractData(file, out)
	count := 0
	for range out {
		count++
	}
	if count != 5 {
		t.Errorf("Data extraction should have outputed 5 elements")
	}
}
