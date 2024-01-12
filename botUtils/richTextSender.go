package botUtils

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/slainsama/msgr_server/globals"
	"io/ioutil"
	"net/http"
)

type Message struct {
	ChatId string `xml:"chat_id"`
	Photo  string `xml:"photo"`
	Text   string `xml:"text"`
}

func SendTextMessage(message *Message) {
	url := globals.UnmarshaledConfig.Bot.Token + "sendmessage"
	params := map[string]string{
		"chat_id": message.ChatId,
		"text":    message.Text,
	}
	reqURL := buildURL(url, params)
	response, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))
}

func SendPhotoMessage(message *Message) {
	url := globals.UnmarshaledConfig.Bot.Token + "sendPhoto"
	params := map[string]string{
		"chat_id": message.ChatId,
		"photo":   message.Photo,
	}
	reqURL := buildURL(url, params)
	response, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))

}

func buildURL(baseURL string, params map[string]string) string {
	url := baseURL + "?"
	for key, value := range params {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	return url
}

func HandleMsg(msg string) *Message {
	decodedMsgBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		fmt.Println("Error decode base64:", err)
		return nil
	}
	msg = string(decodedMsgBytes)
	var message = new(Message)
	err = xml.Unmarshal([]byte(msg), &message)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil
	}
	return message
}
