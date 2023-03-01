package strategies

import (
	"github.com/xgourmandin/slurp"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"github.com/xgourmandin/slurp/internal/handlers/strategies/pagination"
)

func CreatePaginationStrategy(apiConfig slurp.ApiConfiguration, dataStrategy strategies.DataStrategy) strategies.PaginationStrategy {
	switch apiConfig.PaginationConfig.PaginationType {
	case "PAGE_LIMIT":
		return &pagination.PageLimitPaginationStrategy{
			PageParam:     apiConfig.PaginationConfig.PageParam,
			LimitParam:    apiConfig.PaginationConfig.LimitParam,
			CurrentPage:   1,
			LimitValue:    apiConfig.PaginationConfig.PageSize,
			MoreItemsPath: nil,
			DataStrategy:  dataStrategy,
		}
	case "OFFSET_LIMIT":
		return &pagination.OffsetLimitPaginationStrategy{
			OffsetParam:   apiConfig.PaginationConfig.PageParam,
			LimitParam:    apiConfig.PaginationConfig.LimitParam,
			CurrentOffset: 0,
			LimitValue:    apiConfig.PaginationConfig.PageSize,
			DataStrategy:  dataStrategy,
		}
	case "HATEOAS":
		return &pagination.HateoasPaginationStrategy{
			NextLinkPath: apiConfig.PaginationConfig.NextLinkPath,
			DataStrategy: dataStrategy,
		}
	default:
		return pagination.NoPaginationStrategy{}
	}
}
