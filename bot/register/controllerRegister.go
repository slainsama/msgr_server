package register

import (
	"github.com/slainsama/msgr_server/bot"
	"github.com/slainsama/msgr_server/bot/models"
	"github.com/slainsama/msgr_server/utils"
)

func TelegramBotControllerRegister(newUpdate models.TelegramUpdate) {
	commands, args := utils.ExtractCommands(newUpdate)
	for _, command := range commands {
		newHandleUpdate := models.HandleUpdate{NewUpdate: newUpdate, Args: args[command]}
		switch command {
		case "/start":
			utils.RunInGoroutine(newHandleUpdate, bot.startController)
		case "/tasks":
			utils.RunInGoroutine(newHandleUpdate, bot.getUserTaskController)
		case "/addParams":
			utils.RunInGoroutine(newHandleUpdate, bot.addParamsController)
		case "/createTask":
			utils.RunInGoroutine(newHandleUpdate, bot.createUserTaskController)
		}
	}
}
