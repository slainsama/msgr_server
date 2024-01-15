package globals

import "github.com/slainsama/msgr_server/botController"

func botMessageHandler() {
	for {
		select {
		case newUpdate := <-MessageChannel:
			botController.TelegramBotControllerRegister(newUpdate)
		}
	}
}
