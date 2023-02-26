package tests

import (
	"net/http"
	"os"
	"slurp/internal/handlers/strategies"
	"slurp/internal/handlers/strategies/pagination"
	"testing"
)

func TestHateoasPagination(t *testing.T) {
	strategy := pagination.HateoasPaginationStrategy{
		NextLinkPath: "$.next",
		DataStrategy: strategies.JsonDataStrategy{DataRootPath: "$"},
	}

	req, _ := http.NewRequest("GET", "https://test.api.com", nil)
	nextRequest := strategy.ApplyPagination(*req, nil)
	if nextRequest.URL.String() != "https://test.api.com" {
		t.Errorf("First request shall not be tampered")
	}
	pokeabilities, _ := os.ReadFile("./pokemon.json")
	nextRequest = strategy.ApplyPagination(*req, &pokeabilities)
	if nextRequest.URL.String() != "https://pokeapi.co/api/v2/ability?offset=20&limit=20" {
		t.Errorf("Next URL shall be retrieved from the 'next' link")
	}
	if nextRequest.Method != req.Method {
		t.Errorf("Pagination process shall not tamper with HTTP Method")
	}
}
