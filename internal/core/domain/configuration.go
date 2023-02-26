package domain

import (
	"errors"
	"fmt"
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
	OutputConfig          OutputConfig
}

var PaginationType = []string{"NONE", "LIMIT_OFFSET", "PAGE_LIMIT", "TOKEN", "HATEOAS"}

type PaginationConfiguration struct {
	PaginationType string
	PageParam      string
	LimitParam     string
	PageSize       int
	NextLinkPath   string
}

var AuthType = []string{"API_KEY"}

type AuthenticationConfig struct {
	AuthType   string
	InHeader   bool
	TokenEnv   string
	TokenParam string
}

var OutputType = []string{"LOG", "FILE", "BUCKET", "BIGQUERY"}

type OutputConfig struct {
	OutputType string
	FileName   string
	BucketName string
	Project    string
	Dataset    string
	Table      string
	Autodetect bool
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
		if pageType, ok := paginationBlock.(map[string]interface{})["type"]; ok && utils.Contains(PaginationType, pageType.(string)) {
			c.PaginationConfig.PaginationType = pageType.(string)
		} else {
			return errors.New(fmt.Sprintf("No pagination type or wrong value given: %s", pageType))
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
		if nextLinkPath, ok := paginationBlock.(map[string]interface{})["next_link_path"]; ok {
			c.PaginationConfig.NextLinkPath = nextLinkPath.(string)
		}
	} else {
		c.PaginationConfig.PaginationType = "NONE"
	}
	c.AuthConfig = AuthenticationConfig{}
	if authenticationBlock, ok := config["auth"]; ok {
		if authType, ok := authenticationBlock.(map[string]interface{})["type"]; ok && utils.Contains(AuthType, authType.(string)) {
			c.AuthConfig.AuthType = authType.(string)
		} else {
			return errors.New(fmt.Sprintf("No authentication type or wrong value given: %s", authType))
		}
		if inHeader, ok := authenticationBlock.(map[string]interface{})["in_header"]; ok {
			c.AuthConfig.InHeader = inHeader.(bool)
		}
		if tokenEnv, ok := authenticationBlock.(map[string]interface{})["token_env"]; ok {
			c.AuthConfig.TokenEnv = tokenEnv.(string)
		}
		if tokenParam, ok := authenticationBlock.(map[string]interface{})["token_param"]; ok {
			c.AuthConfig.TokenParam = tokenParam.(string)
		}
	} else {
		c.AuthConfig.AuthType = "NONE"
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
	c.OutputConfig = OutputConfig{}
	if outputBlock, ok := config["output"]; ok {
		if outputType, ok := outputBlock.(map[string]interface{})["type"]; ok && utils.Contains(OutputType, outputType.(string)) {
			c.OutputConfig.OutputType = outputType.(string)
		} else {
			return errors.New(fmt.Sprintf("No output type or wrong value given: %s", outputType))
		}
		if fileName, ok := outputBlock.(map[string]interface{})["filename"]; ok {
			c.OutputConfig.FileName = fileName.(string)
		}
		if bucket, ok := outputBlock.(map[string]interface{})["bucket"]; ok {
			c.OutputConfig.BucketName = bucket.(string)
		}
		if project, ok := outputBlock.(map[string]interface{})["project"]; ok {
			c.OutputConfig.Project = project.(string)
		}
		if dataset, ok := outputBlock.(map[string]interface{})["dataset"]; ok {
			c.OutputConfig.Dataset = dataset.(string)
		}
		if table, ok := outputBlock.(map[string]interface{})["table"]; ok {
			c.OutputConfig.Table = table.(string)
		}
		if autodetect, ok := outputBlock.(map[string]interface{})["autodetect"]; ok {
			c.OutputConfig.Autodetect = autodetect.(bool)
		} else {
			c.OutputConfig.Autodetect = false
		}
	} else {
		c.OutputConfig.OutputType = "LOG"
	}
	return nil
}
