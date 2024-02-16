package register

import (
	"github.com/slainsama/msgr_server/bot/controller"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
)

func TelegramBotControllerRegister(newUpdate models.TelegramUpdate) {
	commands, args := botUtils.ExtractCommands(newUpdate)
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
