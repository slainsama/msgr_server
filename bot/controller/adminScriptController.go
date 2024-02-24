package controller

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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
		handler.NewLocalPersistence(),

		handler.StateMap{
			start:        {startHandler},
			uploadScript: {},
			end:          {endHandler},
		},
		[]*handler.Handler{endHandler},
		time.Minute,
		adminAddScriptTimeoutController,
		handler.NewKeyLock(),
	)

	botGlobals.Dispatcher.AddHandler(h)
}

func adminAddScriptTimeoutController(u *types.TelegramUpdate) int {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Timeout")
	return handler.HandleSuccess
}

func adminScriptStartController(u *types.TelegramUpdate) int {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "Please send the script file")
	return uploadScript
}

func adminScriptEndController(u *types.TelegramUpdate) int {
	userID := u.Message.From.ID
	botMethod.SendTextMessage(userID, "process stopped")
	return handler.HandleSuccess
}

func ShouldHandle(u *types.TelegramUpdate) bool {
	return u.Message.Document.FileSize != 0
}

func HandlerFunc(u *types.TelegramUpdate) {
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
