package handler

import (
	"github.com/slainsama/msgr_server/bot/types"
	"github.com/smallnest/safemap"
)

const (
	BufferSize = 1000
)

type UpdateDispatcher struct {
	handlers       []*Handler
	userChannelMap *safemap.SafeMap[PersistenceKey, chan *types.TelegramUpdate]
}

func NewUpdateDispatcher() *UpdateDispatcher {
	return &UpdateDispatcher{
		userChannelMap: safemap.New[PersistenceKey, chan *types.TelegramUpdate](),
	}
}

func (u *UpdateDispatcher) AddHandler(h *Handler) {
	u.handlers = append(u.handlers, h)
}

func (u *UpdateDispatcher) Dispatch(tu *types.TelegramUpdate) {
	pk := PersistenceKey{
		UserID: tu.Message.From.ID,
		ChatID: tu.Message.Chat.ID,
	}
	ch, ok := u.userChannelMap.Get(pk)
	if ok {
		ch <- tu
		return
	} else {
		ch = make(chan *types.TelegramUpdate, BufferSize)
		u.userChannelMap.Set(pk, ch)
		ch <- tu

		go func(ch chan *types.TelegramUpdate) {
			for updateMessage := range ch {
				for _, handler := range u.handlers {
					if handler.Filter(updateMessage) {
						for _, handle := range handler.Handles {
							if result := handle(updateMessage); result == HandleFailed || result == HandleSuccess {
								close(ch)
								u.userChannelMap.Remove(pk)
								return
							}
						}
						break
					}
				}
			}
		}(ch)
	}
}
