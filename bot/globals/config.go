package globals

import "github.com/slainsama/msgr_server/bot/types"

var Config types.BotConfig

// Telegram constants
const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)
const (
	MethodGetFile     = "getFile"
	MethodGetUpdates  = "getUpdates"
	MethodSendMessage = "sendMessage"
	MethodSendPhoto   = "sendPhoto"
)
