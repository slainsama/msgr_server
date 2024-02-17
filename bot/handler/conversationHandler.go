package handler

import (
	"time"

	"github.com/slainsama/msgr_server/models"
)

// ConversationState is a map that represents the state of a conversation
type StateKey struct {
	ConversationID string
	ChatID         int
	UserID         int
}

// conversationState is a map that stores the states of user
var conversationState = make(map[StateKey]int)

// conversationTimeout holds each user's conversation timeout timer
var conversationTimeout = make(map[StateKey]*time.Timer)

type ConversationHandler struct {
	ConversationID string

	StartHandler         *CommandHandler
	ConversationHandlers map[int][]Handler
	EndHandler           *CommandHandler

	ConversationTimeout time.Duration
	TimeoutTask         func()
}

func NewConversationHandler(
	conversationID string,
	startHandler *CommandHandler,
	conversationHandlers map[int][]Handler,
	endHandler *CommandHandler,
) *ConversationHandler {
	return &ConversationHandler{
		ConversationID:       conversationID,
		StartHandler:         startHandler,
		ConversationHandlers: conversationHandlers,
		EndHandler:           endHandler,
		ConversationTimeout:  time.Minute,
	}
}

// SetConversationTimeout sets the timeout for the conversation
func (c *ConversationHandler) SetConversationTimeout(timeout time.Duration) {
	c.ConversationTimeout = timeout
}

func (c *ConversationHandler) SetTimeoutTask(task func()) {
	c.TimeoutTask = task
}

func (c *ConversationHandler) ShouldHandle(u *models.TelegramUpdate) bool {
	_, ok := conversationState[StateKey{
		ConversationID: c.ConversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}]
	return ok || c.StartHandler.ShouldHandle(u) || c.EndHandler.ShouldHandle(u)
}

func (c *ConversationHandler) HandlerFunc(u *models.TelegramUpdate) {
	if c.StartHandler.ShouldHandle(u) {
		c.StartHandler.HandlerFunc(u)

		key := StateKey{
			ConversationID: c.ConversationID,
			ChatID:         u.Message.Chat.ID,
			UserID:         u.Message.From.ID,
		}
		conversationTimeout[key] = time.AfterFunc(c.ConversationTimeout, func() {
			// Execute timeout task
			c.TimeoutTask()
			// Remove the conversation state
			delete(conversationState, key)
		})

		return
	}

	if c.EndHandler.ShouldHandle(u) {
		c.EndHandler.HandlerFunc(u)

		// Clear user's conversation state
		key := StateKey{
			ConversationID: c.ConversationID,
			ChatID:         u.Message.Chat.ID,
			UserID:         u.Message.From.ID,
		}
		delete(conversationState, key)

		// Stop and delete user's conversation timer
		if _, ok := conversationTimeout[key]; ok {
			conversationTimeout[key].Stop()
			delete(conversationTimeout, key)
		}

		return
	}

	state := conversationState[StateKey{
		ConversationID: c.ConversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}]
	for _, handler := range c.ConversationHandlers[state] {
		if handler.ShouldHandle(u) {
			handler.HandlerFunc(u)

			// Reset the user's conversation timeout
			key := StateKey{
				ConversationID: c.ConversationID,
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			}
			if _, ok := conversationTimeout[key]; ok {
				conversationTimeout[key].Reset(c.ConversationTimeout)
			}

			return
		}
	}
}

func UpdateState(state int, key *StateKey) {
	conversationState[*key] = state
}
