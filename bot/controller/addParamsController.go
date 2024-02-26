package controller

import (
	"fmt"

	"github.com/slainsama/msgr_server/bot/botMethod"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/bot/types"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/globals"
)

func init() {
	botGlobals.Dispatcher.AddHandler(handler.NewCommandHandler("/addParams", addParamsController))
}

// addParamsController "/addParams {taskId} {arg1} {arg2}"
func addParamsController(u *types.TelegramUpdate) int {
	userId := u.Message.Chat.ID

	commands, messageArgs := botUtils.ExtractCommands(u)
	taskId := messageArgs[commands[0]][0]
	args := messageArgs[commands[0]][1:]
	//check taskId
	if _, ok := globals.TaskList[taskId]; ok {
		task := globals.TaskList[taskId]
		for _, arg := range args {
			task.Params = append(task.Params, arg)
			botMethod.SendTextMessage(userId, fmt.Sprintf("arg:'%s' added", arg))
		}
	} else {
		botMethod.SendTextMessage(userId, fmt.Sprintf("no %s task", taskId))
		return handler.HandleFailed
	}
	return handler.HandleSuccess
}
