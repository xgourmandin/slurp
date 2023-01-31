package handlers

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"slurp/internal/core/domain"
)

type HttpGetHandler struct {
}

type HttpPostHandler struct {
}

func (HttpGetHandler) SendRequest(ctx domain.Context) []byte {
	req, err := ctx.CreateRequest()
	dump, err := httputil.DumpRequestOut(req, true)
	log.Println(string(dump))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	return respBody
}
