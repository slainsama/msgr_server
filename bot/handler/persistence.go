package handler

import (
	"github.com/smallnest/safemap"
)

// ConversationPersistence interface tells conversation where to store & how to retrieve the current state of the conversation,
// i. e. which "step" the given user is currently at.
type ConversationPersistence interface {
	// GetState & SetState tell conversation handlers how to retrieve & set conversation state.
	GetState(pk PersistenceKey) (int, bool)
	SetState(pk PersistenceKey, state int)
}

// PersistenceKey contains user & chat IDs. It is used to identify conversations with different users in different chats.
type PersistenceKey struct {
	UserID int64
	ChatID int64
}

type LocalPersistence struct {
	States *safemap.SafeMap[PersistenceKey, int]
}

func (p *LocalPersistence) GetState(pk PersistenceKey) (int, bool) {
	return p.States.Get(pk)
}

func (p *LocalPersistence) SetState(pk PersistenceKey, state int) {
	p.States.Set(pk, state)
}

// NewLocalPersistence creates new instance of LocalPersistence.
func NewLocalPersistence() *LocalPersistence {
	return &LocalPersistence{
		States: safemap.New[PersistenceKey, int](),
	}
}
