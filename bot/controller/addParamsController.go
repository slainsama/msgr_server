package controller

import (
	"fmt"

	"github.com/slainsama/msgr_server/bot/botMethod"
	"github.com/slainsama/msgr_server/bot/models"
	"github.com/slainsama/msgr_server/globals"
)

// "/addParams {taskId} {arg1} {arg2}"
func AddParamsController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	taskId := newHandleUpdate.Args[0]
	args := newHandleUpdate.Args[1:]
	//check taskId
	if _, ok := globals.TaskList[taskId]; ok {
		task := globals.TaskList[taskId]
		for _, arg := range args {
			task.Params = append(task.Params, arg)
			botMethod.SendTextMessage(userId, fmt.Sprintf("arg:'%s' added", arg))
		}
	} else {
		botMethod.SendTextMessage(userId, fmt.Sprintf("no %s task", taskId))
	}
}
