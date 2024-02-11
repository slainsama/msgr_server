package botController

import (
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/scriptUtils"
)

func getUserTaskController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	var message string
	for _, task := range globals.TaskList {
		if task.UserId == userId {
			message = message + task.Id + " " + task.ScriptName
		}
	}
	botUtils.SendTextMessage(userId, message)
}

func createUserTaskController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	newTask := scriptUtils.TaskCreate(userId)
}
