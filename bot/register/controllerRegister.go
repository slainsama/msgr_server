package register

import (
	"github.com/slainsama/msgr_server/bot/controller"
	"github.com/slainsama/msgr_server/bot/models"
	"github.com/slainsama/msgr_server/utils"
)

func TelegramBotControllerRegister(newUpdate models.TelegramUpdate) {
	commands, args := utils.ExtractCommands(newUpdate)
	for _, command := range commands {
		newHandleUpdate := models.HandleUpdate{NewUpdate: newUpdate, Args: args[command]}
		switch command {
		case "/start":
			utils.RunInGoroutine(newHandleUpdate, controller.StartController)
		case "/tasks":
			utils.RunInGoroutine(newHandleUpdate, controller.GetUserTaskController)
		case "/addParams":
			utils.RunInGoroutine(newHandleUpdate, controller.AddParamsController)
		case "/createTask":
			utils.RunInGoroutine(newHandleUpdate, controller.CreateUserTaskController)
		}
	}
}
