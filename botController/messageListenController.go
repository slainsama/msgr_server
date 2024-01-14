package botController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"log"
)

func MessageListenController(context *gin.Context) {
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
