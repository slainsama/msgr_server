package botController

import (
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
)

func taskController(newUpdate models.TelegramUpdate) {
	userId := newUpdate.Message.Chat.ID
	var message string
	for _, task := range globals.TaskList {
		if task.UserId == userId {
			message = message + task.Id + " " + task.ScriptName
		}
	}
	botUtils.SendTextMessage(userId, message)
}
