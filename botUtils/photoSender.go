package botUtils

import (
	"log"
	"net/http"

	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/utils"
)

func SendPhotoMessage(chatID int, caption string, photo []byte) {
	config := globals.UnmarshaledConfig

	params := map[string]any{
		"chat_id":    chatID,
		"caption":    caption,
		"parse_mode": "MarkdownV2", // Send as Markdown text
		"photo":      photo,
	}
	url := config.Bot.APIUrl + config.Bot.Token + config.Bot.Methods.SendPhoto
	code, body, err := utils.HttpPOST(url, params)
	if err != nil || code != http.StatusOK {
		log.Println("Error sending PHOTO message:", err)
		log.Println("Response Body:", string(body))
		return
	}
}
