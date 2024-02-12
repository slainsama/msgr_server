package botController

import (
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
)

func TelegramBotControllerRegister(newUpdate models.TelegramUpdate) {
	commands, args := utils.ExtractCommands(newUpdate)
	for _, command := range commands {
		newHandleUpdate := models.HandleUpdate{NewUpdate: newUpdate, Args: args[command]}
		switch command {
		case "/start":
			utils.RunInGoroutine(newHandleUpdate, startController)
		case "/tasks":
			utils.RunInGoroutine(newHandleUpdate, getUserTaskController)
		case "/addParams":
			utils.RunInGoroutine(newHandleUpdate, addParamsController)
		case "/createTask":
			utils.RunInGoroutine(newHandleUpdate, createUserTaskController)
		}
	}
}
