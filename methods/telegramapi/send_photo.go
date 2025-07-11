package telegramapi

import (
	"encoding/json"
	"fmt"
	"resty.dev/v3"
	"strconv"
)

const (
	_sendPhotoMethod = "sendPhoto"
)

func (c *TelegramApi) SendPhoto(request SendPhotoRequest) (result SendMessageResponse, err error) {
	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	keyboardJSON, err := json.Marshal(request.InlineKeyboardMarkup)
	if err != nil {
		err = fmt.Errorf("error keyboard: %v", err)
		return
	}
	keyboardString := string(keyboardJSON)
	if request.InlineKeyboardMarkup == nil {
		keyboardString = ""
	}

	_, err = client.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetMultipartFormData(map[string]string{
			"chat_id":      strconv.FormatInt(request.ChatId, 10),
			"caption":      request.Caption,
			"parse_mode":   request.ParseMode,
			"reply_markup": keyboardString,
		}).
		SetFile("photo", request.Photo).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", c.url, _sendPhotoMethod))
	if err != nil {
		return
	}

	//fmt.Println(res.String())

	return
}
