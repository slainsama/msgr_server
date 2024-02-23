package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/types"
	"github.com/slainsama/msgr_server/utils"
)

var lastUpdateId = 0

func initLastUpdateID() {
	// fetch the latest update message
	baseURL := fmt.Sprintf(globals.APIEndpoint, globals.Config.Token, globals.MethodGetUpdates)
	code, body, err := utils.HttpGET(baseURL, map[string]string{"offset": "-1"})
	if err != nil || code != http.StatusOK {
		log.Fatal(err)
	}

	// marshal the response body
	var messageJson types.TelegramUpdateResponse
	if err := json.Unmarshal(body, &messageJson); err != nil {
		log.Fatal(err)
	}

	// restore the last update message's id
	newUpdates := messageJson.Result
	if len(newUpdates) > 0 {
		lastUpdateId = newUpdates[0].UpdateID
	}
}

func WebhookMessageListenController(context *gin.Context) {
	var messageJson types.TelegramUpdateResponse
	err := context.ShouldBind(&messageJson)
	if err != nil {
		log.Println(err)
	}
	if !messageJson.OK {
		log.Println("err webhook.")
		context.Abort()
	}
	newUpdates := messageJson.Result
	for _, update := range newUpdates {
		messageChannel <- update //消息入队
	}
}

func requestMessageListenController() {
	url := fmt.Sprintf(globals.APIEndpoint, globals.Config.Token, globals.MethodGetUpdates)
	code, body, err := utils.HttpGET(url, nil)
	if err != nil || code != http.StatusOK {
		log.Println(err)
		return
	}

	// marshal the response body
	var messageJson types.TelegramUpdateResponse
	if err := json.Unmarshal(body, &messageJson); err != nil {
		log.Fatal(err)
	}

	newUpdates := messageJson.Result
	for _, update := range newUpdates {
		if update.UpdateID > lastUpdateId {
			messageChannel <- update //消息入队
			lastUpdateId = update.UpdateID
		}
	}
}

func InitRequestMessageListenController() {
	initLastUpdateID() // Get the latest update message's ID
	for {
		requestMessageListenController()
		time.Sleep(5 * time.Second)
	}
}
