package app

import (
	"github.com/slainsama/msgr_server/bot/botMethod"
	"github.com/slainsama/msgr_server/bot/controller"
	"github.com/slainsama/msgr_server/bot/handler"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
)

func botMessageHandler() {
	dispatcher := handler.NewUpdateDispatcher()

	// Register dispatcher's handler
	dispatcher.AddHandler(handler.NewCommandHandler("/start", func(u *models.TelegramUpdate) {
		commands, args := botUtils.ExtractCommands(u)
		hu := models.HandleUpdate{NewUpdate: *u, Args: args[commands[0]]}
		controller.StartController(hu)
	}))
	dispatcher.AddHandler(handler.NewCommandHandler("/tasks", func(u *models.TelegramUpdate) {
		commands, args := botUtils.ExtractCommands(u)
		hu := models.HandleUpdate{NewUpdate: *u, Args: args[commands[0]]}
		controller.GetUserTaskController(hu)
	}))
	dispatcher.AddHandler(handler.NewCommandHandler("/addParams", func(u *models.TelegramUpdate) {
		commands, args := botUtils.ExtractCommands(u)
		hu := models.HandleUpdate{NewUpdate: *u, Args: args[commands[0]]}
		controller.AddParamsController(hu)
	}))
	dispatcher.AddHandler(handler.NewCommandHandler("/createTask", func(u *models.TelegramUpdate) {
		commands, args := botUtils.ExtractCommands(u)
		hu := models.HandleUpdate{NewUpdate: *u, Args: args[commands[0]]}
		controller.CreateUserTaskController(hu)
	}))

	const (
		startUploadScript = iota
		uploadScript
		endUploadScript
	)
	dispatcher.AddHandler(handler.NewConversationHandler(
		"upload_script",
		handler.NewCommandHandler("/start_upload_script", func(u *models.TelegramUpdate) {
			botMethod.SendTextMessage(u.Message.Chat.ID, "Start uploading")
			handler.UpdateState(uploadScript, &handler.StateKey{
				ConversationID: "upload_script",
				ChatID:         u.Message.Chat.ID,
				UserID:         u.Message.From.ID,
			})
		}),
		map[int][]handler.Handler{
			uploadScript: {
				handler.NewCommandHandler("/upload_script", func(u *models.TelegramUpdate) {
					botMethod.SendTextMessage(u.Message.Chat.ID, "Upload the script")
					handler.UpdateState(endUploadScript, &handler.StateKey{
						ConversationID: "upload_script",
						ChatID:         u.Message.Chat.ID,
						UserID:         u.Message.From.ID,
					})
				}),
			},
		},
		handler.NewCommandHandler("/stop_upload_script", func(u *models.TelegramUpdate) {
			botMethod.SendTextMessage(u.Message.Chat.ID, "End uploading script")
		}),
	))

	for {
		newUpdate := <-globals.MessageChannel
		dispatcher.Dispatch(&newUpdate)
	}
}

func InitBot() {
	go botMessageHandler()
}
