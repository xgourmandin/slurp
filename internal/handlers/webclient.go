package handlers

import (
	"fmt"
	"github.com/xgourmandin/slurp/internal/core/ports"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

type HttpHandler struct {
}

func (HttpHandler) SendRequest(ctx ports.Context) []byte {
	req, err := ctx.CreateRequest()
	dump, err := httputil.DumpRequestOut(req, true)
	log.Println(string(dump))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("error closing the response body: %v", err))
		}
	}(resp.Body)
	respBody, _ := io.ReadAll(resp.Body)
	return respBody
}
