package botUtils

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"io"
	"log"
	"net/http"
)

func SendTextMessage(message *models.Message) {
	url := globals.UnmarshaledConfig.Bot.Token + "sendMessage"
	params := map[string]string{
		"chat_id": message.ChatId,
		"text":    message.Text,
	}
	reqURL := buildURL(url, params)
	response, err := http.Get(reqURL)
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing :", err)
		}
	}(response.Body)
	var buf bytes.Buffer
	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("Response Body:", buf.String())
}

func SendPhotoMessage(message *models.Message) {
	url := globals.UnmarshaledConfig.Bot.Token + "sendPhoto"
	params := map[string]string{
		"chat_id": message.ChatId,
		"photo":   message.Photo,
	}
	reqURL := buildURL(url, params)
	response, err := http.Get(reqURL)
	if err != nil {
		log.Println("Error sending GET request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing :", err)
		}
	}(response.Body)
	var buf bytes.Buffer
	_, err = io.Copy(&buf, response.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	log.Println("Response Body:", buf.String())
}

func buildURL(baseURL string, params map[string]string) string {
	url := baseURL + "?"
	for key, value := range params {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	return url
}

func HandleMsg(msg string) *models.Message {
	decodedMsgBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		log.Println("Error decode base64:", err)
		return nil
	}
	msg = string(decodedMsgBytes)
	var message = new(models.Message)
	err = xml.Unmarshal([]byte(msg), &message)
	if err != nil {
		log.Println("Error unmarshalling XML:", err)
		return nil
	}
	return message
}
