package telegramapi

import (
	"fmt"
	"resty.dev/v3"
)

const (
	_answerCallbackQueryMethod = "answerCallbackQuery"
)

func (c *TelegramApi) AnswerCallbackQuery(request AnswerCallbackQueryRequest) error {
	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	res, err := client.R().
		SetBody(request).
		Post(fmt.Sprintf("%s/%s", c.url, _answerCallbackQueryMethod))
	if err != nil {
		return err
	}

	fmt.Println(res.String())

	return nil
}
