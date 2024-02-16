package handler

import "github.com/slainsama/msgr_server/models"

// ConversationState is a map that represents the state of a conversation
type StateKey struct {
	ConversationID string
	ChatID         int
	UserID         int
}

// conversationState is a map that stores the states of user
var conversationState = make(map[StateKey]int)

type ConversationHandler struct {
	ConversationID string

	StartHandler         *CommandHandler
	ConversationHandlers map[int][]Handler
	EndHandler           *CommandHandler
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
	}
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
		return
	}

	if c.EndHandler.ShouldHandle(u) {
		c.EndHandler.HandlerFunc(u)
		// Remove the conversation state
		delete(conversationState, StateKey{
			ConversationID: c.ConversationID,
			ChatID:         u.Message.Chat.ID,
			UserID:         u.Message.From.ID,
		})
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
			return
		}
	}
}

func UpdateState(state int, key *StateKey) {
	conversationState[*key] = state
}
