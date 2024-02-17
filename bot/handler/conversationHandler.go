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

type HandlerMap map[int][]Handler

type ConversationHandler struct {
	conversationID string

	startHandler         *CommandHandler
	conversationHandlers HandlerMap
	endHandler           *CommandHandler

	conversationTimeout time.Duration
	timeoutTask         func(u *models.TelegramUpdate)
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
		timeoutTask:          func(u *models.TelegramUpdate) {},
	}
}

// SetConversationTimeout sets the timeout for the conversation
func (c *ConversationHandler) SetConversationTimeout(timeout time.Duration) {
	c.conversationTimeout = timeout
}

func (c *ConversationHandler) SetTimeoutTask(task func(u *models.TelegramUpdate)) {
	c.timeoutTask = task
}

func (c *ConversationHandler) ShouldHandle(u *models.TelegramUpdate) bool {
	_, ok := conversationState[StateKey{
		ConversationID: c.conversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}]
	if c.startHandler.ShouldHandle(u) {
		return !ok
	} else {
		return ok
	}
}

func (c *ConversationHandler) HandlerFunc(u *models.TelegramUpdate) {
	if c.startHandler.ShouldHandle(u) {
		c.startHandler.HandlerFunc(u)

		key := StateKey{
			ConversationID: c.conversationID,
			ChatID:         u.Message.Chat.ID,
			UserID:         u.Message.From.ID,
		}

		taskWrapper := func(u *models.TelegramUpdate) func() {
			// Create a new task
			return func() {
				c.timeoutTask(u)
				delete((conversationState), key)
			}
		}
		task := taskWrapper(u)
		conversationTimeout[key] = time.AfterFunc(c.conversationTimeout, task)

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

func UpdateState(conversationID string, state int, u *models.TelegramUpdate) {
	key := StateKey{
		ConversationID: conversationID,
		ChatID:         u.Message.Chat.ID,
		UserID:         u.Message.From.ID,
	}
	conversationState[key] = state
}
