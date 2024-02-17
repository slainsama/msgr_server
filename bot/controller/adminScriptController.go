package controller

import (
	"github.com/slainsama/msgr_server/bot/botMethod"
	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/models"
)

const (
	start = iota
	uploadScript
	end
)

func init() {
	startHandler := handler.NewCommandHandler("/admin_upload_script", adminScriptStartController)
	uploadHandler := handler.NewCommandHandler("/upload", adminAddScriptController)
	endHandler := handler.NewCommandHandler("/admin_end_script", adminScriptEndController)

	h := handler.NewConversationHandler(
		"admin_add_script",
		startHandler,
		handler.HandlerMap{uploadScript: {uploadHandler}},
		endHandler,
	)
	h.SetTimeoutTask(adminAddScriptTimeoutController)

	globals.Dispatcher.AddHandler(h)
}

func adminAddScriptTimeoutController(u *models.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Timeout")
}

func adminScriptStartController(u *models.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Please send the script name")
	handler.UpdateState("admin_add_script", uploadScript, u)
}

func adminAddScriptController(u *models.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "test test")
	handler.UpdateState("admin_add_script", end, u)
}

func adminScriptEndController(u *models.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "process stopped")
}
