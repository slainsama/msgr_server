package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/bot/app"
	"github.com/slainsama/msgr_server/globals"
)

func initBotListener(botListener *gin.RouterGroup) {
	if globals.UnmarshaledConfig.Bot.GetUpdate == "webhook" {
		botListener.POST("webhookListener", app.WebhookMessageListenController)
	}
	if globals.UnmarshaledConfig.Bot.GetUpdate == "request" {
		go app.InitRequestMessageListenController()
	}
}
