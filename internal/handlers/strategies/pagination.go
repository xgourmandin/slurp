package strategies

import (
	"slurp/internal/core/ports"
	"slurp/internal/core/ports/strategies"
	"slurp/internal/handlers/strategies/pagination"
)

func CreatePaginationStrategy(apiConfig ports.ApiConfiguration, dataStrategy strategies.DataStrategy) strategies.PaginationStrategy {
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
