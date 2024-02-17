package app

import (
	_ "github.com/slainsama/msgr_server/bot/controller"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/globals"
)

func botMessageHandler() {
	for {
		newUpdate := <-globals.MessageChannel
		botGlobals.Dispatcher.Dispatch(&newUpdate)
	}
}

func InitBot() {
	go botMessageHandler()
}
