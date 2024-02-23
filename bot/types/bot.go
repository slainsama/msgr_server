package types

type Bot struct {
	ID     int
	Token  string
	Secret string
}

type Message struct {
	ChatId int    `xml:"chat_id"`
	Data   string `xml:"text"`
}

type TelegramResponse struct {
	OK        bool `json:"ok"`
	ErrorCode int  `json:"error_code"`
}

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
		Document struct {
			FileName     string `json:"file_name"`
			MimeType     string `json:"mime_type"`
			FileID       string `json:"file_id"`
			FileUniqueID string `json:"file_unique_id"`
			FileSize     int    `json:"file_size"`
		} `json:"document"`
	} `json:"message"`
}

// TelegramUpdateResponse is the response from the GetUpdates method
type TelegramUpdateResponse struct {
	TelegramResponse
	Result []TelegramUpdate `json:"result"`
}

type TelegramFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

type TelegramGetFileResponse struct {
	TelegramResponse
	Result TelegramFile `json:"result"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type HandleUpdate struct {
	NewUpdate TelegramUpdate
	Args      []string
}
