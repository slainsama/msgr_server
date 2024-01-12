package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/userController"
)

func loadUserRouter(userRouter *gin.RouterGroup) {
	userRouter.GET("/:token/getAllScripts", userController.GetAllScripts)
	userRouter.GET("/:token/getAllTasks", userController.GetAllTasks)
}
