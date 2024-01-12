package server

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/userController"
)

// 这里需要加个中间件鉴权
func loadAdminUserRouter(userRouter *gin.RouterGroup) {
	adminRouter := userRouter.Group("/admin")
	adminRouter.GET("/getAllScripts", userController.GetAllScripts)
	adminRouter.GET("/getAllTasks", userController.GetAllTasks)
}
