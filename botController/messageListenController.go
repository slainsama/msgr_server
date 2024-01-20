package botController

import (
	"encoding/json"
	"io"
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
	body, err := utils.HttpGET(baseURL, map[string]string{"offset": "-1"})
	if err != nil {
		log.Fatal(err)
	}

	// marshal the response body
	var messageJson models.TelegramResponse
	if err := json.Unmarshal(body, &messageJson); err != nil {
		log.Fatal(err)
	}

	// restore the last update message's id
	newUpdates := messageJson.Result
	lastUpdateId = newUpdates[0].UpdateID
}

func WebhookMessageListenController(context *gin.Context) {
	var messageJson models.TelegramResponse
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
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: %s\n", resp.Status)
	}
	var messageJson models.TelegramResponse
	err = json.NewDecoder(resp.Body).Decode(&messageJson)
	if err != nil {
		log.Println(err)
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
