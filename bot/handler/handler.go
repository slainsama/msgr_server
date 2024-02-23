package handler

import (
	"github.com/slainsama/msgr_server/bot/types"
)

type Handler interface {
	ShouldHandle(u *types.TelegramUpdate) bool
	HandlerFunc(u *types.TelegramUpdate)
}
