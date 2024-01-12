package userController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/utils"
)

func GetAllScripts(context *gin.Context) {
	//调用昊哥sdk
}

func GetAllTasks(context *gin.Context) {
	var tasksNow []gin.H
	for _, task := range globals.TaskList {
		tasksNow = append(tasksNow, gin.H{task.Id: task.ScriptName})
	}
	utils.SuccessResp(context, gin.H{"tasks": tasksNow})
}
