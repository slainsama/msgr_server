package botMethod

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/types"
	"github.com/slainsama/msgr_server/utils"
)

func GetFile(fileId string) *types.TelegramFile {
	params := map[string]string{
		"file_id": fileId,
	}
	url := fmt.Sprintf(globals.APIEndpoint, globals.Config.Token, globals.MethodGetFile)
	code, body, err := utils.HttpGET(url, params)
	if err != nil || code != http.StatusOK {
		log.Println("Error getting file:", err)
		log.Println("Response Body:", string(body))
		return nil
	}

	// Unmarshal the response
	var fileResp types.TelegramGetFileResponse
	if err := json.Unmarshal(body, &fileResp); err != nil {
		log.Println("Error unmarshaling file response:", err)
		return nil
	} else {
		return &fileResp.Result
	}
}
