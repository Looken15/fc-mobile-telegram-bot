package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
)

func (s TelegramService) Response(params models.TelegramUpdate) (err error) {
	err = s.telegramApi.SendMessage(telegramapi.SendMessageRequest{
		ChatId: params.Message.Chat.ID,
		Text:   params.Message.Text,
	})
	if err != nil {
		return err
	}

	return
}
