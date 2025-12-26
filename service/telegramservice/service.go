package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/utils"
)

type TelegramService struct {
	telegramApi *telegramapi.TelegramApi

	messageId    int64
	userId       int64
	chatId       int64
	username     string
	callbackData utils.CallbackData
}

func New(telegramApi *telegramapi.TelegramApi) TelegramService {
	return TelegramService{
		telegramApi: telegramApi,
	}
}
