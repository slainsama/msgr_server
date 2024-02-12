package scriptController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/utils"
)

// GetParamsController give one param per request
func GetParamsController(context *gin.Context) {
	taskId := context.Param("taskID")
	selectedTask := globals.TaskList[taskId]
	param := selectedTask.Params[0]
	utils.SuccessResp(context, param)
	selectedTask.Params = selectedTask.Params[1:]
}
