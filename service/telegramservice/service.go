package telegramservice

import "fc-mobile-telegram-bot/methods/telegramapi"

type TelegramService struct {
	telegramApi *telegramapi.TelegramApi
}

func New(telegramApi *telegramapi.TelegramApi) TelegramService {
	return TelegramService{
		telegramApi: telegramApi,
	}
}
