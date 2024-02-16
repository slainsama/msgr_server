package handler

import (
	"github.com/slainsama/msgr_server/models"
)

type Handler interface {
	ShouldHandle(u *models.TelegramUpdate) bool
	HandlerFunc(u *models.TelegramUpdate)
}

var Handlers []Handler

func AddHandler(handler Handler) {
	Handlers = append(Handlers, handler)
}
