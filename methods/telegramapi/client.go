package telegramapi

import (
	"fmt"
	"resty.dev/v3"
)

type TelegramApi struct {
	url string
}

func New(url string) *TelegramApi {
	return &TelegramApi{
		url: url,
	}
}

func (c *TelegramApi) SendMessage(request SendMessageRequest) error {
	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	_, err := client.R().
		SetBody(request).
		Post(fmt.Sprintf("%s/sendMessage", c.url))
	if err != nil {
		return err
	}

	return nil
}
