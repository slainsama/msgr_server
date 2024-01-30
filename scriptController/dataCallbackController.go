package scriptController

import (
	"github.com/gin-gonic/gin"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
	"log"
	"mime/multipart"
)

func DataCallbackController(context *gin.Context) {
	taskId := context.Param("taskId")
	task := globals.TaskList[taskId]
	var data models.Callback
	err := context.ShouldBind(&data)
	if err != nil {
		log.Println(err)
		return
	}
	file, _, err := context.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	data.File = file
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(file)
	task.CallbackData <- data
	utils.SuccessResp(context, nil)
}
