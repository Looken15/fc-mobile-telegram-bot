package models

type TelegramUpdate struct {
	EditedMessage *Message       `json:"edited_message"`
	Message       *Message       `json:"message"`
	UpdateID      float64        `json:"update_id"`
	CallbackQuery *CallbackQuery `json:"callback_query"`
}

type CallbackQuery struct {
	Id             string   `json:"id"`
	From           User     `json:"from"`
	Message        *Message `json:"message"`
	Data           string   `json:"data"`
	ChatInstanceId int64    `json:"chat_instance"`
}

type Message struct {
	Chat      Chat   `json:"chat"`
	Date      int64  `json:"date"`
	From      User   `json:"from"`
	MessageID int64  `json:"message_id"`
	Text      string `json:"text"`
}

type Chat struct {
	FirstName string `json:"first_name"`
	ID        int64  `json:"id"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
	Username  string `json:"username"`
}

type User struct {
	FirstName    string `json:"first_name"`
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	LanguageCode string `json:"language_code"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
}
