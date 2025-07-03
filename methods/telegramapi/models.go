package telegramapi

import (
	"fc-mobile-telegram-bot/models"
)

type SendMessageRequest struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
	//ReplyKeyboardMarkup  ReplyKeyboardMarkup  `json:"reply_markup"`
	InlineKeyboardMarkup InlineKeyboardMarkup `json:"reply_markup"`
	ParseMode            string               `json:"parse_mode"`
}

type SendMessageResponse struct {
	Ok     bool           `json:"ok"`
	Result models.Message `json:"result"`
}

type SendPhotoRequest struct {
	ChatId               int64                 `json:"chat_id"`
	Caption              string                `json:"caption"`
	InlineKeyboardMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	ParseMode            string                `json:"parse_mode"`
	Photo                string                `json:"photo"`
}

type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard bool               `json:"resize_keyboard"`
}

type KeyboardButton struct {
	Text string `json:"text"`
}

type DeleteMessageRequest struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
}

type GetChatMemberRequest struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

type GetChatMemberResponse struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	User   User   `json:"user"`
	Status string `json:"status"`
}

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
