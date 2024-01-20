package globals

import (
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/models"
)

func scriptCallbackHandler() {
	for {
		for _, task := range TaskList {
			select {
			case newCallback := <-task.CallbackData:
				switch newCallback.Action {
				case "sendText":
					sendText(task, newCallback)
				}
			}
		}
	}
}

func sendText(task models.Task, newCallback models.Callback) {
	var message models.Message
	message.Data = newCallback.Data.(string)
	message.ChatId = task.UserId
	botUtils.SendTextMessage(message)
}

func initScriptCallback() {
	go scriptCallbackHandler()
}
