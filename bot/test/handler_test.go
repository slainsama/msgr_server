package test

import (
	"sync"
	"testing"
	"time"

	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/bot/types"
)

func TestNewConversationHandler(t *testing.T) {
	updateChan := make(chan types.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	dispatcher := handler.NewUpdateDispatcher()

	// Initialize conversation handler
	conversationHandler := handler.NewConversationHandler(
		handler.NewLocalPersistence(),
		handler.StateMap{
			start: {
				handler.NewCommandHandler("/startUpload", func(u *types.TelegramUpdate) int {
					t.Log("Start upload")
					return sendHello
				}),
			},
			sendHello: {
				handler.NewCommandHandler("/hello", func(u *types.TelegramUpdate) int {
					t.Log("Hello")
					return end
				}),
			},
			end: {
				handler.NewCommandHandler("/endUpload", func(u *types.TelegramUpdate) int {
					t.Log("End upload")
					return handler.HandleSuccess
				}),
			},
		},
		[]*handler.Handler{},

		time.Minute,
		nil,
	)
	dispatcher.AddHandler(conversationHandler)

	// Mock update message channel
	wg.Add(2)
	go func() {
		// Send start message
		updateChan <- newStartUpdate()
		// Send message
		updateChan <- newHelloUpdate()
		// Send end message
		updateChan <- newEndUpdate()

		wg.Done()
	}()

	go func() {
		for i := 0; i < 3; i++ {
			newUpdate := <-updateChan
			dispatcher.Dispatch(&newUpdate)
		}
		wg.Done()
	}()
	time.Sleep(time.Second * 1)

	wg.Wait()
}

func TestNewConversationHandlerWithMultiChoice(t *testing.T) {
	updateChan := make(chan types.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)

	dispatcher := handler.NewUpdateDispatcher()

	// Initialize conversation handler
	conversationHandler := handler.NewConversationHandler(
		handler.NewLocalPersistence(),
		handler.StateMap{
			start: {
				handler.NewCommandHandler("/startUpload", func(u *types.TelegramUpdate) int {
					t.Log("Start upload")
					return sendHello
				}),
			},
			sendHello: {
				handler.NewCommandHandler("/hello", func(u *types.TelegramUpdate) int {
					t.Log("Hello")
					return end
				}),
				handler.NewCommandHandler("/another_hello", func(u *types.TelegramUpdate) int {
					t.Log("Another Hello")
					return end
				}),
			},
			end: {
				handler.NewCommandHandler("/endUpload", func(u *types.TelegramUpdate) int {
					t.Log("End upload")
					return handler.HandleFailed
				}),
			},
		},
		[]*handler.Handler{},

		time.Minute,
		nil,
	)
	dispatcher.AddHandler(conversationHandler)

	// Mock update message channel
	wg.Add(2)
	go func() {
		// Send start message
		updateChan <- newStartUpdate()
		// Send message
		updateChan <- newAnotherHelloUpdate()
		// Send end message
		updateChan <- newEndUpdate()
		wg.Done()
	}()

	go func() {
		for i := 0; i < 3; i++ {
			newUpdate := <-updateChan
			dispatcher.Dispatch(&newUpdate)
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	wg.Wait()
}

func TestNewConversationHandlerWithoutPermission(t *testing.T) {
	updateChan := make(chan types.TelegramUpdate, 5)
	var wg sync.WaitGroup
	const (
		start = iota
		sendHello
		end
	)
	role := make(map[int64]bool)
	role[123] = true
	role[426] = false

	dispatcher := handler.NewUpdateDispatcher()

	// Initialize conversation handler
	conversationHandler := handler.NewConversationHandler(
		handler.NewLocalPersistence(),
		handler.StateMap{
			start: {
				handler.NewCommandHandler("/startUpload", func(u *types.TelegramUpdate) int {
					userID := u.Message.From.ID
					t.Log(userID, " Start upload")
					return sendHello
				}),
			},
			sendHello: {
				handler.NewCommandHandler("/hello", func(u *types.TelegramUpdate) int {
					userID := u.Message.From.ID
					if role[userID] {
						t.Log(userID, " Pass Auth")
						return end
					} else {
						t.Log(userID, " No permission")
						return handler.HandleFailed
					}
				}),
			},
			end: {
				handler.NewCommandHandler("/endUpload", func(u *types.TelegramUpdate) int {
					userID := u.Message.From.ID
					t.Log(userID, " End upload")
					return handler.HandleFailed
				}),
			},
		},
		[]*handler.Handler{},

		time.Minute,
		nil,
	)
	dispatcher.AddHandler(conversationHandler)

	// Mock update message channel
	wg.Add(2)
	go func() {
		updateChan <- newStartWithoutPermissionUpdate()
		updateChan <- newStartUpdate()
		updateChan <- newHelloWithoutPermissionUpdate()
		updateChan <- newHelloUpdate()
		updateChan <- newEndWithoutPermissionUpdate()
		updateChan <- newEndUpdate()
		wg.Done()
	}()

	go func() {
		for i := 0; i < 6; i++ {
			newUpdate := <-updateChan
			dispatcher.Dispatch(&newUpdate)
		}
		wg.Done()
	}()

	time.Sleep(time.Second)

	wg.Wait()
}

// func TestConversationHandlerWithTimeout(t *testing.T) {
// 	updateChan := make(chan types.TelegramUpdate, 5)
// 	var wg sync.WaitGroup
// 	const (
// 		start = iota
// 		sendHello
// 		end
// 	)

// 	dispatcher := handler.NewUpdateDispatcher()

// 	// Initialize conversation handler
// 	conversationHandler := handler.NewConversationHandler(
// 		"testUpload",
// 		handler.NewCommandHandler("/startUpload", func(u *types.TelegramUpdate) {
// 			t.Log("Start upload")
// 			handler.UpdateState("testUpload", sendHello, u)
// 		}),
// 		handler.HandlerMap{
// 			sendHello: {
// 				handler.NewCommandHandler("/hello", func(u *types.TelegramUpdate) {
// 					t.Log("Hello")
// 					handler.UpdateState("testUpload", end, u)
// 				}),
// 			},
// 		},
// 		handler.NewCommandHandler("/endUpload", func(u *types.TelegramUpdate) {
// 			t.Log("End upload")
// 		}),
// 	)
// 	conversationHandler.SetConversationTimeout(time.Second)
// 	conversationHandler.SetTimeoutTask(func(u *types.TelegramUpdate) { t.Log("Timeout", u.Message.From.ID) })

// 	dispatcher.AddHandler(conversationHandler)

// 	// Mock update message channel
// 	wg.Add(2)
// 	go func() {
// 		// Send start message
// 		updateChan <- newStartUpdate()
// 		time.Sleep(time.Second * 2)
// 		// Send message
// 		updateChan <- newHelloUpdate()
// 		time.Sleep(time.Second)
// 		wg.Done()
// 	}()

// 	go func() {
// 		for i := 0; i < 2; i++ {
// 			newUpdate := <-updateChan
// 			dispatcher.Dispatch(&newUpdate)
// 		}
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }
