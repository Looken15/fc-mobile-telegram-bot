package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
	"fmt"
)

const (
	_lastUpdateDate = "2 февраля, 2025"

	_sendPhotoCaption = "ТОП-10 %s в FC Mobile\n\nПоследнее обновление:\n%s"
)

func (s TelegramService) Response(params models.TelegramUpdate) (err error) {
	if params.Message.Text == "Вратари" {
		err = s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:  params.Message.Chat.ID,
			Photo:   fmt.Sprintf("./images/%s.jpg", "gk"),
			Caption: fmt.Sprintf(_sendPhotoCaption, "Вратарей", _lastUpdateDate),
		})
		if err != nil {
			return err
		}

		return nil
	}

	err = s.telegramApi.SendMessage(telegramapi.SendMessageRequest{
		ChatId: params.Message.Chat.ID,
		Text:   params.Message.Text,
	})
	if err != nil {
		return err
	}

	return
}
