package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/botController"
)

func initBotListener(botListener *gin.RouterGroup) {
	botListener.POST("webhookListener", botController.MessageListenController)
}
