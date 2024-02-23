package botMethod

import (
	"fmt"
	"log"
	"net/http"

	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/utils"
)

func SendPhotoMessage(chatID int, caption string, photo []byte) {

	params := map[string]any{
		"chat_id":    chatID,
		"caption":    caption,
		"parse_mode": "MarkdownV2", // Send as Markdown text
		"photo":      photo,
	}
	url := fmt.Sprintf(globals.APIEndpoint, globals.Config.Token, globals.MethodSendPhoto)
	code, body, err := utils.HttpPOST(url, params)
	if err != nil || code != http.StatusOK {
		log.Println("Error sending PHOTO message:", err)
		log.Println("Response Body:", string(body))
		return
	}
}
