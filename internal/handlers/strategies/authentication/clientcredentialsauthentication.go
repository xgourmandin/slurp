package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

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
		clientSecret, err := s.SecretManager.GetSecretValue(s.ClientSecret)
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
		log.Printf("Authenticating with URL %s\n", s.AccessTokenUrl)
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
			log.Printf("Error with the authentication request response. Status of authentication request is %d\n", resp.StatusCode)
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
	log.Println("Authentication succeeded")
	return req
}
