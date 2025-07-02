package telegramapi

import (
	"fmt"
	"resty.dev/v3"
	"strconv"
)

const (
	_sendPhotoMethod = "sendPhoto"
)

func (c *TelegramApi) SendPhoto(request SendPhotoRequest) error {
	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	res, err := client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("photo", request.Photo).
		SetFormData(map[string]string{
			"chat_id": strconv.FormatInt(request.ChatId, 10)}).
		Post(fmt.Sprintf("%s/%s", c.url, _sendPhotoMethod))
	if err != nil {
		return err
	}

	fmt.Println(res.String())

	return nil
}
