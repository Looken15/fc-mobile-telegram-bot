package telegramapi

type TelegramApi struct {
	url string
}

func New(url string) *TelegramApi {
	return &TelegramApi{
		url: url,
	}
}
