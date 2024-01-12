package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/scriptController"
)

func loadScriptRouter(scriptRouter *gin.RouterGroup) {
	scriptRouter.GET("/:secretKey/:taskId/params", scriptController.GetParamsController)
	scriptRouter.POST("/:secretKey/:taskId/sendData", scriptController.DataCallbackController)
}
