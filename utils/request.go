package utils

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/slainsama/msgr_server/globals"
)

// HttpGET is a wrapper of http.Get
// It returns a *http.Response
func HttpGET(baseURL string, params map[string]string) (code int, body []byte, err error) {
	url := buildURL(baseURL, params)

	if isDEBUG := globals.UnmarshaledConfig.DEBUG.Switch; isDEBUG {
		log.Println("HttpGET:", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Return response body as []byte
	body, err = io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

// HttpPOST is a wrapper of http.Post which support sending photos
func HttpPOST(baseURL string, params map[string]any) (code int, body []byte, err error) {
	if isDEBUG := globals.UnmarshaledConfig.DEBUG.Switch; isDEBUG {
		log.Println("HttpPOST:", baseURL)
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	for key, value := range params {
		switch v := value.(type) {
		case string:
			_ = writer.WriteField(key, v)
		case []byte:
			imgReader := bytes.NewReader(v)
			part, err := writer.CreateFormFile(key, "photo")
			if err != nil {
				return 0, nil, err
			}
			// 将图片文件内容复制到字段中
			_, err = io.Copy(part, imgReader)
			if err != nil {
				return 0, nil, err
			}
		default:
			// Handle other types if necessary
		}
	}

	err = writer.Close()
	if err != nil {
		return 0, nil, err
	}

	resp, err := http.Post(baseURL, writer.FormDataContentType(), &requestBody)
	if err != nil {
		return 0, nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Return response body as []byte
	body, err = io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

// buildURL builds an url with params
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
