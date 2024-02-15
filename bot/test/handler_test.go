package test

import (
	"sync"
	"testing"
	"time"

	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/models"
)

func TestNewCoversationHandler(t *testing.T) {
	updateChan := make(chan models.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	// Initialize conversation handler
	conversationHandler := handler.NewCoversationHandler(
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
		update := models.TelegramUpdate{}
		update.Message.Text = "/startUpload"
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 12,
			Type:   "bot_command",
		})
		updateChan <- update
		time.Sleep(time.Second)

		// Send message
		update = models.TelegramUpdate{}
		update.Message.Text = "/hello"
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 6,
			Type:   "bot_command",
		})
		updateChan <- update
		time.Sleep(time.Second)

		// Send end message
		update = models.TelegramUpdate{}
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Text = "/endUpload"
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 10,
			Type:   "bot_command",
		})
		updateChan <- update
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

func TestNewCoversationHandlerWithMultiChoice(t *testing.T) {
	updateChan := make(chan models.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	// Initialize conversation handler
	conversationHandler := handler.NewCoversationHandler(
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
		update := models.TelegramUpdate{}
		update.Message.Text = "/startUpload"
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 12,
			Type:   "bot_command",
		})
		updateChan <- update
		time.Sleep(time.Second)

		// Send message
		update = models.TelegramUpdate{}
		update.Message.Text = "/another_hello"
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 14,
			Type:   "bot_command",
		})
		updateChan <- update
		time.Sleep(time.Second)

		// Send end message
		update = models.TelegramUpdate{}
		update.Message.Chat.ID = 123
		update.Message.From.ID = 123
		update.Message.Text = "/endUpload"
		update.Message.Entities = append(update.Message.Entities, struct {
			Offset int    "json:\"offset\""
			Length int    "json:\"length\""
			Type   string "json:\"type\""
		}{
			Offset: 0,
			Length: 10,
			Type:   "bot_command",
		})
		updateChan <- update
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
