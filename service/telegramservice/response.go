package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
	"fmt"
	"github.com/samber/lo"
)

const (
	_startMessage  = "/start"
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
	if params.CallbackQuery != nil && lo.Contains(_positionsArray, params.CallbackQuery.Data) {
		position := params.CallbackQuery.Data

		err = s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:    params.CallbackQuery.Message.Chat.ID,
			Caption:   fmt.Sprintf(_sendPhotoCaption, _positionsWordMap[position], _lastUpdateDate),
			ParseMode: _htmlParseMode,
			Photo:     fmt.Sprintf(_imagePath, position),
		})
		if err != nil {
			return err
		}

		return nil
	}

	if params.Message != nil && params.Message.Text == _startMessage {
		keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
		for _, pos := range _positionsArray {
			keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
			keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: pos, CallbackData: pos})

			keyboard = append(keyboard, keyboardArray)
		}

		err = s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
			ChatId:               params.Message.Chat.ID,
			Caption:              fmt.Sprintf(_helloCaption, params.Message.From.Username),
			InlineKeyboardMarkup: telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
			ParseMode:            _htmlParseMode,
			Photo:                fmt.Sprintf(_imagePath, "hello"),
		})
		if err != nil {
			return err
		}

		return nil
	}

	return
}
