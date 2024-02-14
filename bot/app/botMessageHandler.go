package app

import (
	"github.com/slainsama/msgr_server/bot/register"
	"github.com/slainsama/msgr_server/globals"
)

func botMessageHandler() {
	for {
		select {
		case newUpdate := <-globals.MessageChannel:
			register.TelegramBotControllerRegister(newUpdate)
		}
	}
}

func InitBot() {
	go botMessageHandler()
}
