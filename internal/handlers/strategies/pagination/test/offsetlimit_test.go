package tests

import (
	"github.com/xgourmandin/slurp/internal/handlers/strategies/pagination"
	"net/http"
	"testing"
)

func TestOffsetLimitPaginationUpdate(t *testing.T) {
	strategy := pagination.OffsetLimitPaginationStrategy{
		OffsetParam:   "offset",
		LimitParam:    "per_page",
		CurrentOffset: 0,
		LimitValue:    25,
		DataStrategy:  nil,
	}

	req, _ := http.NewRequest("GET", "https://test.api.com", nil)
	paginated := strategy.ApplyPagination(*req, nil)
	if paginated.URL.Query().Get("offset") != "0" {
		t.Errorf("Offset parameter is not correctly set on first call")
	}
	paginated = strategy.ApplyPagination(paginated, nil)
	if paginated.URL.Query().Get("offset") != "25" {
		t.Errorf("Offset parameter is not correctly set on second call")
	}
}
