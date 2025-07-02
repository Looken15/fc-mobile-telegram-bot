package telegramapi

type SendMessageRequest struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type SendPhotoRequest struct {
	ChatId int64  `json:"chat_id"`
	Photo  string `json:"photo"`
}
