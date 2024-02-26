package types

type Bot struct {
	ID     int
	Token  string
	Secret string
}

type TelegramResponse struct {
	OK        bool `json:"ok"`
	ErrorCode int  `json:"error_code"`
}

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	// optional
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int       `json:"message_id"`
	From      *User     `json:"from,omitempty"`
	Chat      *Chat     `json:"sender_chat,omitempty"`
	Date      int       `json:"date"`
	Text      string    `json:"text"`
	Entities  []Entity  `json:"entities,omitempty"`
	Document  *Document `json:"document,omitempty"`
}

// User represents a Telegram user or bot.
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot,omitempty"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	UserName                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

// Chat represents a chat.
type Chat struct {
	ID                    int64  `json:"id"`
	Type                  string `json:"type"`
	Title                 string `json:"title,omitempty"`
	UserName              string `json:"username,omitempty"`
	FirstName             string `json:"first_name,omitempty"`
	LastName              string `json:"last_name,omitempty"`
	Bio                   string `json:"bio,omitempty"`
	HasPrivateForwards    bool   `json:"has_private_forwards,omitempty"`
	Description           string `json:"description,omitempty"`
	InviteLink            string `json:"invite_link,omitempty"`
	SlowModeDelay         int    `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int    `json:"message_auto_delete_time,omitempty"`
	HasProtectedContent   bool   `json:"has_protected_content,omitempty"`
	StickerSetName        string `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool   `json:"can_set_sticker_set,omitempty"`
	LinkedChatID          int64  `json:"linked_chat_id,omitempty"`
}

// Document represents a general file.
type Document struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileName     string `json:"file_name,omitempty"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
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
