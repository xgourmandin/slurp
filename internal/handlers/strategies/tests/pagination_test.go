package tests

import (
	"net/http"
	"slurp/internal/handlers/strategies"
	"testing"
)

func TestPaginationUpdate(t *testing.T) {
	strategy := strategies.PageLimitPaginationStrategy{
		PageParam:     "page",
		LimitParam:    "per_page",
		CurrentPage:   1,
		LimitValue:    25,
		MoreItemsPath: nil,
	}

	req, _ := http.NewRequest("GET", "https://test.api.com", nil)
	paginated := strategy.ApplyPagination(*req)
	if paginated.URL.Query().Get("page") != "1" {
		t.Errorf("Page parameter is not correctly set on first call")
	}
	paginated = strategy.ApplyPagination(paginated)
	if paginated.URL.Query().Get("page") != "2" {
		t.Errorf("Page parameter is not correctly set on second call")
	}
}
