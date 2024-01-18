package botController

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"io"
	"log"
	"net/http"
	"time"
)

var lastUpdateId = 0

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
	url := "https://api.telegram.org/bot" + globals.UnmarshaledConfig.Bot.Token + "/getUpdates"
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
	for {
		requestMessageListenController()
		time.Sleep(5 * time.Second)
	}
}
