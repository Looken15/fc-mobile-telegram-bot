package models

type TelegramUpdate struct {
	EditedMessage *Message `json:"edited_message"`
	Message       *Message `json:"message"`
	UpdateID      float64  `json:"update_id"`
}

type Message struct {
	Chat      Chat    `json:"chat"`
	Date      float64 `json:"date"`
	From      User    `json:"from"`
	MessageID float64 `json:"message_id"`
	Text      string  `json:"text"`
}

type Chat struct {
	FirstName string  `json:"first_name"`
	ID        float64 `json:"id"`
	LastName  string  `json:"last_name"`
	Type      string  `json:"type"`
	Username  string  `json:"username"`
}

type User struct {
	FirstName    string  `json:"first_name"`
	ID           float64 `json:"id"`
	IsBot        bool    `json:"is_bot"`
	LanguageCode string  `json:"language_code"`
	LastName     string  `json:"last_name"`
	Username     string  `json:"username"`
}
