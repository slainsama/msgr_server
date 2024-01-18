package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/botController"
	"github.com/slainsama/msgr_server/globals"
)

func initBotListener(botListener *gin.RouterGroup) {
	if globals.UnmarshaledConfig.Bot.GetUpdate == "webhook" {
		botListener.POST("webhookListener", botController.WebhookMessageListenController)
	}
	if globals.UnmarshaledConfig.Bot.GetUpdate == "request" {
		go botController.InitRequestMessageListenController()
	}
}
