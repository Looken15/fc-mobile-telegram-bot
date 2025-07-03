package telegramapi

import (
	"fmt"
	"resty.dev/v3"
)

const (
	_getChatMemberMethod = "getChatMember"

	_chatId       = -1001805786440
	_chatUsername = "@KaramaFC"
)

func (c *TelegramApi) GetChatMember(request GetChatMemberRequest) (result GetChatMemberResponse, err error) {
	request.ChatId = _chatId

	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	res, err := client.R().
		SetBody(request).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", c.url, _getChatMemberMethod))
	if err != nil {
		return
	}

	fmt.Println(res.String())

	return
}
