package botMethod

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/utils"
)

// SendTextMessage 待修改
func SendTextMessage(chatId int64, data string) {
	url := fmt.Sprintf(globals.APIEndpoint, globals.Config.Token, globals.MethodSendMessage)
	params := map[string]string{
		"chat_id":    strconv.FormatInt(chatId, 10),
		"text":       data,
		"parse_mode": "MarkdownV2", // Send as Markdown text
	}
	code, body, err := utils.HttpGET(url, params)
	if err != nil || code != http.StatusOK {
		log.Println("Error sending TEXT message:", err)
		log.Println("Response Body:", string(body))
		return
	}
}
