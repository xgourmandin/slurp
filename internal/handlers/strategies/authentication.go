package strategies

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"github.com/xgourmandin/slurp/configuration"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"github.com/xgourmandin/slurp/internal/core/ports/strategies"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NoAuthenticationStrategy struct {
}

func (NoAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	return req
}

type ApiTokenAuthenticationStrategy struct {
	SecretManager ports.SecretManager
	Token         string // Not the token itself, but the env variable name that holds the token
	InHeader      bool   // If true, the token goes in an Authentication header, else, it goes in the query string
	AuthParam     string
}

func (s ApiTokenAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	tokenValue, err := s.SecretManager.GetSecretValue(s.Token)
	if err != nil {
		fmt.Printf("%v\n", err)
		return req
	}
	if s.InHeader {
		req.Header.Add(s.AuthParam, tokenValue)
	} else {
		q := req.URL.Query()
		q.Set(s.AuthParam, tokenValue)
		req.URL.RawQuery = q.Encode()
	}
	return req
}

type ClientCredentialsAuthenticationStrategy struct {
	SecretManager   ports.SecretManager
	AccessTokenUrl  string
	PayloadTemplate string
	ClientId        string
	ClientSecret    string
	AccessTokenPath string
	currentToken    *string
}

func (s *ClientCredentialsAuthenticationStrategy) AddAuthentication(req http.Request) http.Request {
	if s.currentToken == nil {
		clientId, err := s.SecretManager.GetSecretValue(s.ClientId)
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}
		clientSecret, err := s.SecretManager.GetSecretValue(s.ClientId)
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}
		payload := strings.Replace(s.PayloadTemplate, "${CLIENT_ID}", clientId, 1)
		payload = strings.Replace(payload, "${CLIENT_SECRET}", clientSecret, 1)
		tokens := strings.Split(payload, "&")
		data := url.Values{}
		for _, token := range tokens {
			keyval := strings.Split(token, "=")
			data.Set(keyval[0], keyval[1])
		}
		tokenReq, err := http.NewRequest("POST", s.AccessTokenUrl, strings.NewReader(data.Encode()))
		tokenReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}
		client := &http.Client{}
		resp, err := client.Do(tokenReq)
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}
		defer resp.Body.Close()

		jsonData := interface{}(nil)
		respBody, err := io.ReadAll(resp.Body)
		err = json.Unmarshal(respBody, &jsonData)
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}
		accessToken, err := jsonpath.Get(s.AccessTokenPath, jsonData)
		if err != nil {
			fmt.Printf("%v\n", err)
			return req
		}

		accessStr := accessToken.(string)
		s.currentToken = &accessStr
	}
	token := *s.currentToken
	req.Header.Add("Authorization", "Bearer "+token)

	return req
}

func CreateAuthenticationStrategy(apiConfig configuration.ApiConfiguration, manager ports.SecretManager) strategies.AuthenticationStrategy {
	switch apiConfig.AuthConfig.AuthType {
	case "API_KEY":
		return ApiTokenAuthenticationStrategy{
			SecretManager: manager,
			Token:         apiConfig.AuthConfig.TokenSecret,
			InHeader:      apiConfig.AuthConfig.InHeader,
			AuthParam:     apiConfig.AuthConfig.TokenParam,
		}
	case "CLIENT_CREDS":
		return &ClientCredentialsAuthenticationStrategy{
			SecretManager:   manager,
			AccessTokenUrl:  apiConfig.AuthConfig.AccessTokenUrl,
			PayloadTemplate: apiConfig.AuthConfig.PayloadTemplate,
			ClientId:        apiConfig.AuthConfig.ClientIdSecret,
			ClientSecret:    apiConfig.AuthConfig.ClientSecretSecret,
			AccessTokenPath: apiConfig.AuthConfig.AccessTokenPath,
		}
	default:
		return NoAuthenticationStrategy{}
	}
}
