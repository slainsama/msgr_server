package botController

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
)

var lastUpdateId = 0

func initLastUpdateID() {
	config := globals.UnmarshaledConfig

	// fetch the latest update message
	baseURL := config.Bot.APIUrl + config.Bot.Token + config.Bot.Methods.GetUpdates
	code, body, err := utils.HttpGET(baseURL, map[string]string{"offset": "-1"})
	if err != nil || code != http.StatusOK {
		log.Fatal(err)
	}

	// marshal the response body
	var messageJson models.TelegramUpdateResponse
	if err := json.Unmarshal(body, &messageJson); err != nil {
		log.Fatal(err)
	}

	// restore the last update message's id
	newUpdates := messageJson.Result
	lastUpdateId = newUpdates[0].UpdateID
}

func WebhookMessageListenController(context *gin.Context) {
	var messageJson models.TelegramUpdateResponse
	err := context.ShouldBind(&messageJson)
	if err != nil {
		log.Println(err)
	}
	if messageJson.OK != true {
		log.Println("err webhook.")
		context.Abort()
	}
	newUpdates := messageJson.Result
	for _, update := range newUpdates {
		globals.MessageChannel <- update //消息入队
	}
}

func requestMessageListenController() {
	config := globals.UnmarshaledConfig

	url := config.Bot.APIUrl + config.Bot.Token + config.Bot.Methods.GetUpdates
	code, body, err := utils.HttpGET(url, nil)
	if err != nil || code != http.StatusOK {
		log.Println(err)
		return
	}

	// marshal the response body
	var messageJson models.TelegramUpdateResponse
	if err := json.Unmarshal(body, &messageJson); err != nil {
		log.Fatal(err)
	}

	newUpdates := messageJson.Result
	for _, update := range newUpdates {
		if update.UpdateID > lastUpdateId {
			globals.MessageChannel <- update //消息入队
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
