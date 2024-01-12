package scriptController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
	"log"
)

func DataCallbackController(context *gin.Context) {
	taskId := context.Param("taskId")
	task := globals.TaskList[taskId]
	var data models.DataModel
	err := context.ShouldBind(&data)
	if err != nil {
		log.Println(err)
	}
	task.CallbackData = append(task.CallbackData, data)
	utils.SuccessResp(context, nil)
}
