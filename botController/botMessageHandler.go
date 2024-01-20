package botController

import "github.com/slainsama/msgr_server/globals"

func botMessageHandler() {
	for {
		select {
		case newUpdate := <-globals.MessageChannel:
			TelegramBotControllerRegister(newUpdate)
		}
	}
}

func initBot() {
	go botMessageHandler()
}
