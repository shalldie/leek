package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type FetchOptions struct {
	Method  string
	Headers map[string]string
	Query   map[string]string
	Params  map[string]any
}

func Fetch(fetchUrl string, fetchOptions *FetchOptions) []byte {
	if fetchOptions == nil {
		fetchOptions = &FetchOptions{}
	}
	client := &http.Client{}

	queryValues := url.Values{}

	for key, val := range fetchOptions.Query {
		queryValues.Add(key, val)
	}

	requrl, _ := url.ParseRequestURI(fetchUrl)
	if fetchOptions.Query != nil {
		requrl.RawQuery = queryValues.Encode()
	}

	var paramBody io.Reader

	if fetchOptions.Params != nil {
		postJSON, _ := json.Marshal(fetchOptions.Params)
		paramBody = strings.NewReader(string(postJSON))
	}

	req, _ := http.NewRequest(fetchOptions.Method, requrl.String(), paramBody)
	for key, val := range fetchOptions.Headers {
		req.Header.Add(key, val)
	}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body
}
