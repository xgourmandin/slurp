package domain

import (
	"errors"
	"slurp/internal/core/utils"
)

var HttpMethod = []string{"GET", "POST"}

var DataType = []string{"JSON", "XML"}

type ApiConfiguration struct {
	Url                   string
	Method                string
	AuthConfig            AuthenticationConfig
	PaginationConfig      PaginationConfiguration
	DataType              string
	DataRoot              string
	AdditionalHeaders     map[string]string
	AdditionalQueryParams map[string]string
}

var PaginationType = []string{"NONE", "LIMIT_OFFSET", "PAGE_LIMIT", "TOKEN", "LINK"}

type PaginationConfiguration struct {
	PaginationType string
	PageParam      string
	LimitParam     string
	PageSize       int
}

type AuthenticationConfig struct {
}

func (c *ApiConfiguration) FromMap(config map[string]interface{}) error {
	if url, ok := config["url"]; ok {
		c.Url = url.(string)
	} else {
		return errors.New("missing url value in API config")
	}
	if method, ok := config["method"]; ok {
		if utils.Contains(HttpMethod, method.(string)) {
			c.Method = method.(string)
		} else {
			return errors.New("Wrong HTTP method given: " + method.(string))
		}
	} else {
		return errors.New("missing http method value in API config")
	}
	c.PaginationConfig = PaginationConfiguration{}
	if paginationBlock, ok := config["pagination"]; ok {
		if pageType, ok := paginationBlock.(map[string]interface{})["type"]; ok {
			c.PaginationConfig.PaginationType = pageType.(string)
		}
		if pageParam, ok := paginationBlock.(map[string]interface{})["page_param"]; ok {
			c.PaginationConfig.PageParam = pageParam.(string)
		}
		if limitParam, ok := paginationBlock.(map[string]interface{})["limit_param"]; ok {
			c.PaginationConfig.LimitParam = limitParam.(string)
		}
		if pageSize, ok := paginationBlock.(map[string]interface{})["page_size"]; ok {
			c.PaginationConfig.PageSize = pageSize.(int)
		}
	} else {
		c.PaginationConfig.PaginationType = "NONE"
	}
	if dataBlock, ok := config["data"]; ok {
		dataType, ok := dataBlock.(map[string]interface{})["type"]
		if ok && utils.Contains(DataType, dataType.(string)) {
			c.DataType = dataType.(string)
		}
		if dataRoot, ok := dataBlock.(map[string]interface{})["root"]; ok {
			c.DataRoot = dataRoot.(string)
		}
	}
	return nil
}
