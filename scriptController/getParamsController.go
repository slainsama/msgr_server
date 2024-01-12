package scriptController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/utils"
)

func GetParamsController(context *gin.Context) {
	taskId := context.Param("taskID")
	selectedTask := globals.TaskList[taskId]
	params := selectedTask.Params
	utils.SuccessResp(context, params)
}
