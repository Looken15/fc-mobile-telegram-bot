package telegramapi

type SendMessageRequest struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
	//ReplyKeyboardMarkup  ReplyKeyboardMarkup  `json:"reply_markup"`
	InlineKeyboardMarkup InlineKeyboardMarkup `json:"reply_markup"`
	ParseMode            string               `json:"parse_mode"`
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
