package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
	"fc-mobile-telegram-bot/utils"
	"fmt"
	"github.com/samber/lo"
)

const (
	_startMessage    = "/start"
	_backMessage     = "/back"
	_tryAgainMessage = "/tryAgain"

	_htmlParseMode = "html"
	_imagePath     = "./images/%s.jpg"

	_lastUpdateDate = "2 февраля, 2025"

	_sendPhotoCaption = "<b>ТОП-10 %s в FC Mobile</b>\n\nПоследнее обновление:\n%s<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 25</a>"

	_helloCaption = "<b>Приветствую, %s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 25</a>"
)

var (
	_membersArray     = []string{"member", "creator", "administrator"}
	_positionsArray   = []string{"ВРТ", "ЛЗ", "ЦЗ", "ПЗ", "ЦОП", "ЛП", "ЦП", "ПП", "ЦАП", "ЛВ", "НАП", "ПВ"}
	_positionsWordMap = map[string]string{
		"ВРТ": "Вратарей",
		"ЛЗ":  "Левых защитников",
		"ЦЗ":  "Центральных защитников",
		"ПЗ":  "Правых защитников",
		"ЦОП": "Центральных опорных полузащитников",
		"ЛП":  "Левых полузащитников",
		"ЦП":  "Центральных полузащитников",
		"ПП":  "Правых полузащитников",
		"ЦАП": "Центральных атакующих полузащитников",
		"ЛВ":  "Левых вингеров",
		"НАП": "Центральных нападающих",
		"ПВ":  "Правых вингеров",
	}
)

func (s TelegramService) Response(params models.TelegramUpdate) (err error) {
	var messageId, userId, chatId int64
	var username string
	if params.Message != nil {
		userId = params.Message.From.ID
		messageId = params.Message.MessageID
		chatId = params.Message.Chat.ID
		username = params.Message.From.Username
	}
	if params.CallbackQuery != nil {
		userId = params.CallbackQuery.From.ID
		messageId = params.CallbackQuery.Message.MessageID
		chatId = params.CallbackQuery.Message.Chat.ID
		username = params.CallbackQuery.From.Username
	}
	var callbackData utils.CallbackData
	if params.CallbackQuery != nil {
		callbackData, err = utils.DecodeCallbackData(params.CallbackQuery.Data)
		if err != nil {
			return err
		}
	}

	result, err := s.telegramApi.GetChatMember(telegramapi.GetChatMemberRequest{
		UserId: userId,
	})
	if err != nil {
		return
	}

	if !lo.Contains(_membersArray, result.Result.Status) {
		keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
		keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

		newCallbackData := utils.CallbackData{
			NextCommand: _tryAgainMessage,
			MessageId:   messageId,
		}
		keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
			Text:         "Проверить подписку",
			CallbackData: utils.EncodeCallbackData(newCallbackData),
		})
		keyboard = append(keyboard, keyboardLine)

		err = s.telegramApi.SendMessage(telegramapi.SendMessageRequest{
			ChatId:               chatId,
			Text:                 "Надо подписаться бро",
			InlineKeyboardMarkup: telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
			ParseMode:            _htmlParseMode,
		})
		if err != nil {
			return err
		}

		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    chatId,
			MessageId: callbackData.MessageId + 1,
		})
		if err != nil {
			return err
		}

		return
	}

	if params.CallbackQuery != nil && lo.Contains(_positionsArray, callbackData.Position) {
		position := callbackData.Position

		keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
		keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

		newCallbackData := utils.CallbackData{
			NextCommand: _backMessage,
			MessageId:   messageId,
		}
		keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: utils.EncodeCallbackData(newCallbackData),
		})
		keyboard = append(keyboard, keyboardLine)

		_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:               chatId,
			Caption:              fmt.Sprintf(_sendPhotoCaption, _positionsWordMap[position], _lastUpdateDate),
			ParseMode:            _htmlParseMode,
			Photo:                fmt.Sprintf(_imagePath, position),
			InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		})
		if err != nil {
			return err
		}

		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    chatId,
			MessageId: callbackData.MessageId + 1,
		})
		if err != nil {
			return err
		}

		return nil
	}

	if callbackData.NextCommand == _tryAgainMessage || callbackData.NextCommand == _backMessage || (params.Message != nil && params.Message.Text == _startMessage) {
		keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
		for _, pos := range _positionsArray {

			keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
			keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: pos, CallbackData: utils.EncodeCallbackData(utils.CallbackData{
				Position:  pos,
				MessageId: messageId,
			})})

			keyboard = append(keyboard, keyboardArray)
		}

		_, err = s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:               chatId,
			Caption:              fmt.Sprintf(_helloCaption, username),
			InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
			ParseMode:            _htmlParseMode,
			Photo:                fmt.Sprintf(_imagePath, "hello"),
		})
		if err != nil {
			return err
		}

		if callbackData.MessageId != 0 {
			err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
				ChatId:    chatId,
				MessageId: callbackData.MessageId + 1,
			})
		}

		return nil
	}

	return
}
