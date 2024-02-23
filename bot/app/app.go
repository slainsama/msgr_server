package app

import (
	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/types"
)

func InitBot(botConfig types.BotConfig) {
	globals.Config = botConfig
	go botMessageHandler()
}
