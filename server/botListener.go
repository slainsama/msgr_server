package server

import "github.com/gin-gonic/gin"

func initBotLisener(botListener *gin.RouterGroup) {
	botListener.POST("webhookListener")
}
