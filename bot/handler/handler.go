package handler

import (
	"strings"
	"time"

	"github.com/slainsama/msgr_server/bot/types"
)

const (
	HandleFailed  = -2
	HandleSuccess = -1
)

// HandleFunc processes update.
type HandleFunc func(u *types.TelegramUpdate) int

// Handler defines a function that will handle updates that pass the filtering.
type Handler struct {
	Filter  FilterFunc
	Handles []HandleFunc
}

// NewHandler creates a new generic handler.
func NewHandler(filter FilterFunc, handles ...HandleFunc) *Handler {
	if filter == nil {
		filter = Any()
	}
	return &Handler{filter, handles}
}

// NewMessageHandler creates a handler for updates that contain message.
func NewMessageHandler(filter FilterFunc, handles ...HandleFunc) *Handler {
	newFilter := IsMessage()
	if filter != nil {
		newFilter = And(newFilter, filter)
	}
	return NewHandler(newFilter, handles...)
}

// NewCommandHandler is an extension for NewMessageHandler that creates a handler for updates that contain message with command.
func NewCommandHandler(command string, handles ...HandleFunc) *Handler {
	handles = append([]HandleFunc{}, handles...)
	commandFilters := []FilterFunc{}
	for _, variant := range strings.Split(command, " ") {
		commandFilters = append(commandFilters, IsCommandMessage(variant))
	}
	newFilter := Or(commandFilters...)
	return NewMessageHandler(
		newFilter,
		handles...,
	)
}

// StateMap is an alias to map of strings to handler slices.
type StateMap map[int][]*Handler

// NewConversationHandler creates a conversation handler.
func NewConversationHandler(
	persistence ConversationPersistence,
	states StateMap,
	cancelHandlers []*Handler, // Handlers that will be used when exit

	timeout time.Duration,
	timeoutTask HandleFunc,
) *Handler {
	handler := &Handler{
		func(u *types.TelegramUpdate) bool {
			user, chat := u.Message.From, u.Message.Chat
			if user == nil || chat == nil {
				return false
			}

			pk := PersistenceKey{user.ID, chat.ID}
			state, ok := persistence.GetState(pk)
			candidates := states[state]
			if ok {
				candidates = append(candidates, cancelHandlers...)
			}

			for _, handler := range candidates {
				if handler.Filter(u) {
					return true // Release the lock in the handles
				}
			}
			return false
		},
		[]HandleFunc{
			func(u *types.TelegramUpdate) int {
				user, chat := u.Message.From, u.Message.Chat
				pk := PersistenceKey{user.ID, chat.ID}

				state, ok := persistence.GetState(pk)
				candidates := states[state]
				if ok {
					candidates = append(candidates, cancelHandlers...)
				}

				var result int
				for _, handler := range candidates {
					if handler.Filter(u) {
						for _, f := range handler.Handles {
							result = f(u)
							if result == HandleFailed {
								break
							} else {
								persistence.SetState(pk, result)
							}
						}
						break
					}
				}
				return result
			},
		},
	}
	return handler
}
