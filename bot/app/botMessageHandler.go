package app

import (
	_ "github.com/slainsama/msgr_server/bot/controller"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
)

func botMessageHandler() {
	for {
		newUpdate := <-messageChannel
		botGlobals.Dispatcher.Dispatch(&newUpdate)
	}
}
