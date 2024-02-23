package app

import "github.com/slainsama/msgr_server/bot/types"

var messageChannel = make(chan types.TelegramUpdate)
