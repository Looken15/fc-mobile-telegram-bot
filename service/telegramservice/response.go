package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
	"fc-mobile-telegram-bot/utils"
	"fmt"
	"github.com/samber/lo"
)

const (
	_startMessage = "/start"
	_backMessage  = "/back"

	_htmlParseMode = "html"
	_imagePath     = "./images/%s.jpg"

	_lastUpdateDate = "2 февраля, 2025"

	_sendPhotoCaption = "<b>ТОП-10 %s в FC Mobile</b>\n\nПоследнее обновление:\n%s"

	_helloCaption = "<b>Приветствую, %s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию"
)

var (
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
	var callbackData utils.CallbackData
	if params.CallbackQuery != nil {
		callbackData, err = utils.DecodeCallbackData(params.CallbackQuery.Data)
		if err != nil {
			return err
		}
	}

	if params.CallbackQuery != nil && lo.Contains(_positionsArray, callbackData.Position) {
		position := callbackData.Position

		keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
		keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

		newCallbackData := utils.CallbackData{
			NextCommand: _backMessage,
			MessageId:   params.CallbackQuery.Message.MessageID,
		}
		keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
			Text:         "Назад",
			CallbackData: utils.EncodeCallbackData(newCallbackData),
		})
		keyboard = append(keyboard, keyboardLine)

		_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:               params.CallbackQuery.Message.Chat.ID,
			Caption:              fmt.Sprintf(_sendPhotoCaption, _positionsWordMap[position], _lastUpdateDate),
			ParseMode:            _htmlParseMode,
			Photo:                fmt.Sprintf(_imagePath, position),
			InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		})
		if err != nil {
			return err
		}

		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    params.CallbackQuery.Message.Chat.ID,
			MessageId: callbackData.MessageId + 1,
		})

		return nil
	}

	if callbackData.NextCommand == _backMessage || (params.Message != nil && params.Message.Text == _startMessage) {
		var messageId int64
		var chatId int64
		var username string
		if params.Message != nil {
			messageId = params.Message.MessageID
			chatId = params.Message.Chat.ID
			username = params.Message.From.Username
		}
		if params.CallbackQuery != nil {
			messageId = params.CallbackQuery.Message.MessageID
			chatId = params.CallbackQuery.Message.Chat.ID
			username = params.CallbackQuery.From.Username
		}

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
