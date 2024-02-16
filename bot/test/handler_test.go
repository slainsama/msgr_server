package test

import (
	"sync"
	"testing"
	"time"

	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/models"
)

func TestNewConversationHandler(t *testing.T) {
	updateChan := make(chan models.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	// Initialize conversation handler
	conversationHandler := handler.NewConversationHandler(
		"testUpload",
		handler.NewCommandHandler("/startUpload", func(u *models.TelegramUpdate) {
			t.Log("Start upload")
			handler.UpdateState(sendHello, &handler.StateKey{
				ConversationID: "testUpload",
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			})
		}),
		map[int][]handler.Handler{
			sendHello: {
				handler.NewCommandHandler("/hello", func(u *models.TelegramUpdate) {
					t.Log("Hello")
					handler.UpdateState(sendHello, &handler.StateKey{
						ConversationID: "testUpload",
						ChatID:         u.Message.Chat.ID,
						UserID:         u.Message.From.ID,
					})
				}),
			},
		},
		handler.NewCommandHandler("/endUpload", func(u *models.TelegramUpdate) {
			t.Log("End upload")
			handler.UpdateState(end, &handler.StateKey{
				ConversationID: "testUpload",
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			})
		}),
	)
	handler.AddHandler(conversationHandler)

	// Mock update message channel
	wg.Add(2)
	go func() {
		// Send start message
		updateChan <- newStartUpdate()
		time.Sleep(time.Second)
		// Send message
		updateChan <- newHelloUpdate()
		time.Sleep(time.Second)
		// Send end message
		updateChan <- newEndUpdate()
		time.Sleep(time.Second)

		// Send end message
		wg.Done()
	}()

	go func() {
		for i := 0; i < 3; i++ {
			newUpdate := <-updateChan
			for _, h := range handler.Handlers {
				if h.ShouldHandle(&newUpdate) {
					h.HandlerFunc(&newUpdate)
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
}

func TestNewConversationHandlerWithMultiChoice(t *testing.T) {
	updateChan := make(chan models.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	// Initialize conversation handler
	conversationHandler := handler.NewConversationHandler(
		"testUpload",
		handler.NewCommandHandler("/startUpload", func(u *models.TelegramUpdate) {
			t.Log("Start upload")
			handler.UpdateState(sendHello, &handler.StateKey{
				ConversationID: "testUpload",
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			})
		}),
		map[int][]handler.Handler{
			sendHello: {
				handler.NewCommandHandler("/hello", func(u *models.TelegramUpdate) {
					t.Log("Hello")
					handler.UpdateState(sendHello, &handler.StateKey{
						ConversationID: "testUpload",
						ChatID:         u.Message.Chat.ID,
						UserID:         u.Message.From.ID,
					})
				}),
				handler.NewCommandHandler("/another_hello", func(u *models.TelegramUpdate) {
					t.Log("Another Hello")
					handler.UpdateState(sendHello, &handler.StateKey{
						ConversationID: "testUpload",
						ChatID:         u.Message.Chat.ID,
						UserID:         u.Message.From.ID,
					})
				}),
			},
		},
		handler.NewCommandHandler("/endUpload", func(u *models.TelegramUpdate) {
			t.Log("End upload")
			handler.UpdateState(end, &handler.StateKey{
				ConversationID: "testUpload",
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			})
		}),
	)
	handler.AddHandler(conversationHandler)

	// Mock update message channel
	wg.Add(2)
	go func() {
		// Send start message
		updateChan <- newStartUpdate()
		time.Sleep(time.Second)
		// Send message
		updateChan <- newAnotherHelloUpdate()
		time.Sleep(time.Second)
		// Send end message
		updateChan <- newEndUpdate()
		time.Sleep(time.Second)
		// Send end message
		wg.Done()
	}()

	go func() {
		for i := 0; i < 3; i++ {
			newUpdate := <-updateChan
			for _, h := range handler.Handlers {
				if h.ShouldHandle(&newUpdate) {
					h.HandlerFunc(&newUpdate)
				}
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
