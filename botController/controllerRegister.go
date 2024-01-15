package botController

import (
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
)

func TelegramBotControllerRegister(newUpdate models.TelegramUpdate) {
	switch newUpdate.Message.Text {
	case "/start":
		utils.RunInGoroutine(newUpdate, startController)
	}
}
