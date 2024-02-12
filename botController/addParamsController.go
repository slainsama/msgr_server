package botController

import (
	"fmt"
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
)

// "/addParams {taskId} {arg1} {arg2}"
func addParamsController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	taskId := newHandleUpdate.Args[0]
	args := newHandleUpdate.Args[1:]
	//check taskId
	if _, ok := globals.TaskList[taskId]; ok {
		task := globals.TaskList[taskId]
		for _, arg := range args {
			task.Params = append(task.Params, arg)
			botUtils.SendTextMessage(userId, fmt.Sprintf("arg:'%s' added", arg))
		}
	} else {
		botUtils.SendTextMessage(userId, fmt.Sprintf("no %s task", taskId))
	}
}
