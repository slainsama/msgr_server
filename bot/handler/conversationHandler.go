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
	conversationID string

	startHandler         *CommandHandler
	conversationHandlers map[int][]Handler
	endHandler           *CommandHandler

	conversationTimeout time.Duration
	timeoutTask         func()
}

func NewConversationHandler(
	conversationID string,
	startHandler *CommandHandler,
	conversationHandlers map[int][]Handler,
	endHandler *CommandHandler,
) *ConversationHandler {
	return &ConversationHandler{
		conversationID:       conversationID,
		startHandler:         startHandler,
		conversationHandlers: conversationHandlers,
		endHandler:           endHandler,
		conversationTimeout:  time.Minute,
		timeoutTask:          func() {},
	}
}

// SetConversationTimeout sets the timeout for the conversation
func (c *ConversationHandler) SetConversationTimeout(timeout time.Duration) {
	c.conversationTimeout = timeout
}

func (c *ConversationHandler) SetTimeoutTask(task func()) {
	c.timeoutTask = task
}

func (c *ConversationHandler) ShouldHandle(u *models.TelegramUpdate) bool {
	_, ok := conversationState[StateKey{
		ConversationID: c.conversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}]
	return ok || c.startHandler.ShouldHandle(u) || c.endHandler.ShouldHandle(u)
}

func (c *ConversationHandler) HandlerFunc(u *models.TelegramUpdate) {
	if c.startHandler.ShouldHandle(u) {
		c.startHandler.HandlerFunc(u)

		key := StateKey{
			ConversationID: c.conversationID,
			ChatID:         u.Message.Chat.ID,
			UserID:         u.Message.From.ID,
		}
		conversationTimeout[key] = time.AfterFunc(c.conversationTimeout, func() {
			// Execute timeout task
			c.timeoutTask()
			// Remove the conversation state
			delete(conversationState, key)
		})

		return
	}

	if c.endHandler.ShouldHandle(u) {
		c.endHandler.HandlerFunc(u)

		// Clear user's conversation state
		key := StateKey{
			ConversationID: c.conversationID,
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
		ConversationID: c.conversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}]
	for _, handler := range c.conversationHandlers[state] {
		if handler.ShouldHandle(u) {
			handler.HandlerFunc(u)

			// Reset the user's conversation timeout
			key := StateKey{
				ConversationID: c.conversationID,
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			}
			if _, ok := conversationTimeout[key]; ok {
				conversationTimeout[key].Reset(c.conversationTimeout)
			}

			return
		}
	}
}

func UpdateState(state int, key *StateKey) {
	conversationState[*key] = state
}
