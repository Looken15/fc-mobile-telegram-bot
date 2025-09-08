package telegramapi

import (
	"fmt"
	"github.com/samber/lo"
	"resty.dev/v3"
)

const (
	_getChatMemberMethod = "getChatMember"

	_karamaChatId   = -1001805786440
	_basementChatId = -1001661886071

	_chatUsername = "@KaramaFC"
)

var (
	_membersArray = []string{"member", "creator", "administrator"}
)

func (c *TelegramApi) CheckIfUserSub(userId int64) (bool, error) {
	karamaRes, err := c.GetChatMember(GetChatMemberRequest{UserId: userId, ChatId: _karamaChatId})
	if err != nil {
		return false, err
	}

	if !lo.Contains(_membersArray, karamaRes.Result.Status) {
		return false, nil
	}

	basementRes, err := c.GetChatMember(GetChatMemberRequest{UserId: userId, ChatId: _basementChatId})
	if err != nil {
		return false, err
	}

	if !lo.Contains(_membersArray, basementRes.Result.Status) {
		return false, nil
	}

	return true, nil
}

func (c *TelegramApi) GetChatMember(request GetChatMemberRequest) (result GetChatMemberResponse, err error) {
	client := resty.New()
	defer func(client *resty.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println("Error closing telegram client")
		}
	}(client)

	_, err = client.R().
		SetBody(request).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", c.url, _getChatMemberMethod))
	if err != nil {
		return
	}

	return
}
