package scriptController

import (
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"io"
	"log"
)

func scriptCallbackHandler() {
	for {
		for _, task := range globals.TaskList {
			select {
			case newCallback := <-task.CallbackData:
				switch newCallback.Action {
				case "sendText":
					sendText(task, newCallback)
				case "sendPhoto":
					sendPhoto(task, newCallback)

				}
			}
		}
	}
}

func sendText(task models.Task, newCallback models.Callback) {
	botUtils.SendTextMessage(task.UserId, newCallback.Msg)
}

func sendPhoto(task models.Task, newCallback models.Callback) {
	fileBytes, err := io.ReadAll(newCallback.File)
	if err != nil {
		log.Println(err)
	}
	botUtils.SendPhotoMessage(task.UserId, "", fileBytes)
}

func initScriptCallback() {
	go scriptCallbackHandler()
}
