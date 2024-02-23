package controller

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slainsama/msgr_server/bot/botMethod"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/bot/types"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/utils"
)

const (
	start = iota
	uploadScript
	end
)

func init() {
	startHandler := handler.NewCommandHandler("/admin_upload_script", adminScriptStartController)
	endHandler := handler.NewCommandHandler("/admin_end_script", adminScriptEndController)

	h := handler.NewConversationHandler(
		"admin_add_script",
		startHandler,
		handler.HandlerMap{
			uploadScript: []handler.Handler{
				&AdminUploadScriptHandler{},
			},
		},
		endHandler,
	)
	h.SetTimeoutTask(adminAddScriptTimeoutController)

	botGlobals.Dispatcher.AddHandler(h)
}

func adminAddScriptTimeoutController(u *types.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Timeout")
}

func adminScriptStartController(u *types.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Please send the script file")
	handler.UpdateState("admin_add_script", uploadScript, u)
}

func adminScriptEndController(u *types.TelegramUpdate) {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "process stopped")
}

// Implement the handler for the uploadScript state
type AdminUploadScriptHandler struct {
}

func (h *AdminUploadScriptHandler) ShouldHandle(u *types.TelegramUpdate) bool {
	return u.Message.Document.FileSize != 0
}

func (h *AdminUploadScriptHandler) HandlerFunc(u *types.TelegramUpdate) {
	FileID := u.Message.Document.FileID
	file := botMethod.GetFile(FileID)

	if file != nil {
		fileURL := fmt.Sprintf(botGlobals.FileEndpoint, botGlobals.Config.Token, file.FilePath)
		_, fileBytes, err := utils.HttpGET(fileURL, nil)
		if err != nil {
			log.Println(err)
			sendMsg := "Error getting file, please try again."
			botMethod.SendTextMessage(u.Message.From.ID, botUtils.EscapeChar(sendMsg))
			return
		}

		// Store fileBytes to filesystem
		s := strings.Split(file.FilePath, "/")
		filePath := botGlobals.Config.ScriptPath + "/" + s[len(s)-1]

		if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
			log.Println(err)
			sendMsg := "Error saving file, please try again."
			botMethod.SendTextMessage(u.Message.From.ID, botUtils.EscapeChar(sendMsg))
			return
		}
	} else {
		sendMsg := "Error getting file, please try again."
		botMethod.SendTextMessage(u.Message.From.ID, botUtils.EscapeChar(sendMsg))
	}
}
