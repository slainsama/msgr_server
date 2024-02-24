package handler

import (
	"github.com/slainsama/msgr_server/bot/types"
)

type UpdateDispatcher struct {
	handlers []*Handler
}

func NewUpdateDispatcher() *UpdateDispatcher {
	return &UpdateDispatcher{}
}

func (u *UpdateDispatcher) AddHandler(h *Handler) {
	u.handlers = append(u.handlers, h)
}

func (u *UpdateDispatcher) Dispatch(tu *types.TelegramUpdate) {
	for _, h := range u.handlers {
		if h.Filter(tu) {
			// Process all the handles in the background
			go func(u *types.TelegramUpdate) {
				for _, f := range h.Handles {
					if result := f(tu); result == HandleFailed {
						break
					}
				}
			}(tu)
			break
		}
	}
}
