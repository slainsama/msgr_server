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
	url := BuildURL(baseURL, params)
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

func BuildURL(baseURL string, params map[string]string) string {
	url := baseURL + "?"
	for key, value := range params {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	return url
}
