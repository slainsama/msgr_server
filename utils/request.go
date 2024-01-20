package utils

import (
	"errors"
	"io"
	"log"
	"net/http"
)

// HttpGET is a wrapper of http.Get
// It returns a *http.Response
func HttpGET(baseURL string, params map[string]string) ([]byte, error) {
	url := buildURL(baseURL, params)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err := errors.New("Error: " + resp.Status)
		return nil, err
	}

	// Return response body as []byte
	body, err := io.ReadAll(resp.Body)

	return body, err
}

// buildURL builds a url with params
// If params is nil, it returns baseURL
func buildURL(baseURL string, params map[string]string) string {
	if params == nil {
		return baseURL
	}
	url := baseURL + "?"
	for key, value := range params {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	return url
}
