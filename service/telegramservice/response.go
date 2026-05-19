package telegramservice

import (
	"fc-mobile-telegram-bot/methods/telegramapi"
	"fc-mobile-telegram-bot/models"
	"fc-mobile-telegram-bot/utils"
	"fmt"
	"github.com/samber/lo"
)

const (
	_shopUrl = "https://t.me/whitegamestore_bot?start=agg_0453030c"

	_startMessage       = "/start"
	_toPositionsMessage = "/positions"
	_toTacticsMessage   = "/tactics"
	_tryAgainMessage    = "/tryAgain"

	_htmlParseMode = "html"
	_imagePathJPG  = "./images/%s.jpg"
	_imagePathPNG  = "./images/%s.png"

	_lastUpdateDate = "19 мая, 2026"

	_sendTacticPhotoCaption   = "<b>Лучшая тактика для схемы %s</b>\n\nПоследнее обновление:\n%s\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"
	_sendPositionPhotoCaption = "<b>ТОП-10 %s в FC Mobile</b>\n\nПоследнее обновление:\n%s\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"

	_subscribeNeededCaption = "Чтобы использовать бота, необходимо подписаться на каналы <a href=\"https://t.me/+mf4AwsUOHlBiNDky\"> KARAMA | FC MOBILE 26 | FIFA MOBILE </a> и <a href=\"https://t.me/+rkUjX8CQwYcwMjQy\"> BASEMENT ATHLETIC | FC MOBILE </a> и нажать кнопку «Проверить подписку»"

	_positionCaption = "<b>Приветствую, @%s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"
	_tacticCaption   = "<b>Приветствую, @%s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"
	_helloCaption    = "<b>Приветствую, @%s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"
)

var (
	_mainButtonsArray = []models.Button{
		{
			Text:        "ТОП 10",
			NextCommand: _toPositionsMessage,
		},
		{
			Text: "Магазин",
			Url:  _shopUrl,
		},
		/*{
			Text:        "Тактики",
			NextCommand: _toTacticsMessage,
		}*/}
	_positionsArray   = []string{"ВРТ", "ЛЗ", "ЦЗ", "ПЗ", "ЦОП", "ЛП", "ЦП", "ПП", "ЦАП", "ЛВ", "НАП", "ПВ"}
	_tacticsArray     = []string{"3-5-2", "3-4-3 (в линию)", "3-4-3 (ромб)", "4-3-3 (атака)", "4-3-3 (удержание)", "4-2-4", "4-2-4 (2)", "4-1-2-1-2 (узкая)", "4-2-2-2", "4-2-2-2 (2)", "4-2-3-1"}
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

func (s *TelegramService) fillParams(userId int64, messageId int64, chatId int64, username string) {
	s.userId = userId
	s.messageId = messageId
	s.chatId = chatId
	s.username = username
}

func (s *TelegramService) fillCallbackData(callbackData string) error {
	result, err := utils.DecodeCallbackData(callbackData)
	if err != nil {
		return err
	}
	s.callbackData = result
	return nil
}

func (s *TelegramService) sendUnsubMessage() error {
	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
	keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

	newCallbackDataCheck := utils.CallbackData{
		NextCommand: _tryAgainMessage,
		MessageId:   s.messageId,
	}
	keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
		Text:         "Проверить подписку",
		CallbackData: utils.EncodeCallbackData(newCallbackDataCheck),
	})
	keyboard = append(keyboard, keyboardLine)

	err := s.telegramApi.SendMessage(telegramapi.SendMessageRequest{
		ChatId:               s.chatId,
		Text:                 _subscribeNeededCaption,
		InlineKeyboardMarkup: telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		ParseMode:            _htmlParseMode,
	})
	if err != nil {
		return err
	}

	err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
		ChatId:    s.chatId,
		MessageId: s.messageId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *TelegramService) sendStartMessage() error {
	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
	for _, button := range _mainButtonsArray {
		keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
		keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: button.Text, Url: button.Url, CallbackData: utils.EncodeCallbackData(utils.CallbackData{
			MessageId:   s.messageId,
			NextCommand: button.NextCommand,
		})})

		keyboard = append(keyboard, keyboardArray)
	}

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_helloCaption, s.username),
		InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		ParseMode:            _htmlParseMode,
		Photo:                fmt.Sprintf(_imagePathJPG, "hello"),
	})
	if err != nil {
		return err
	}

	if s.callbackData.MessageId != 0 {
		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    s.chatId,
			MessageId: s.messageId,
		})
	}

	return nil
}

func (s *TelegramService) sendPositionsPickMessage() error {
	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)

	keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
	keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: "В главное меню", CallbackData: utils.EncodeCallbackData(utils.CallbackData{
		MessageId:   s.messageId,
		NextCommand: _startMessage,
	})})

	keyboard = append(keyboard, keyboardArray)

	for _, pos := range _positionsArray {

		keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
		keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: pos, CallbackData: utils.EncodeCallbackData(utils.CallbackData{
			Position:  pos,
			MessageId: s.messageId,
		})})

		keyboard = append(keyboard, keyboardArray)
	}

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_positionCaption, s.username),
		InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		ParseMode:            _htmlParseMode,
		Photo:                fmt.Sprintf(_imagePathJPG, "hello"),
	})
	if err != nil {
		return err
	}

	if s.callbackData.MessageId != 0 {
		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    s.chatId,
			MessageId: s.messageId,
		})
	}

	return nil
}

func (s *TelegramService) sendPositionsMessage() error {
	position := s.callbackData.Position

	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
	keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

	newCallbackData := utils.CallbackData{
		NextCommand: _toPositionsMessage,
		MessageId:   s.messageId,
	}
	keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: utils.EncodeCallbackData(newCallbackData),
	})
	keyboard = append(keyboard, keyboardLine)

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_sendPositionPhotoCaption, _positionsWordMap[position], _lastUpdateDate),
		ParseMode:            _htmlParseMode,
		Photo:                fmt.Sprintf(_imagePathPNG, position),
		InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
	})
	if err != nil {
		return err
	}

	err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
		ChatId:    s.chatId,
		MessageId: s.messageId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *TelegramService) sendTacticsPickMessage() error {
	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)

	keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
	keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: "В главное меню", CallbackData: utils.EncodeCallbackData(utils.CallbackData{
		MessageId:   s.messageId,
		NextCommand: _startMessage,
	})})

	keyboard = append(keyboard, keyboardArray)

	for _, tactic := range _tacticsArray {

		keyboardArray := make([]telegramapi.InlineKeyboardButton, 0)
		keyboardArray = append(keyboardArray, telegramapi.InlineKeyboardButton{Text: tactic, CallbackData: utils.EncodeCallbackData(utils.CallbackData{
			Tactic:    tactic,
			MessageId: s.messageId,
		})})

		keyboard = append(keyboard, keyboardArray)
	}

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_tacticCaption, s.username),
		InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
		ParseMode:            _htmlParseMode,
		Photo:                fmt.Sprintf(_imagePathJPG, "hello"),
	})
	if err != nil {
		return err
	}

	if s.callbackData.MessageId != 0 {
		err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
			ChatId:    s.chatId,
			MessageId: s.messageId,
		})
	}

	return nil
}

func (s *TelegramService) sendTacticMessage() error {
	tactic := s.callbackData.Tactic

	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
	keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

	newCallbackData := utils.CallbackData{
		NextCommand: _toTacticsMessage,
		MessageId:   s.messageId,
	}
	keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: utils.EncodeCallbackData(newCallbackData),
	})
	keyboard = append(keyboard, keyboardLine)

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_sendTacticPhotoCaption, tactic, _lastUpdateDate),
		ParseMode:            _htmlParseMode,
		Photo:                fmt.Sprintf(_imagePathPNG, tactic),
		InlineKeyboardMarkup: &telegramapi.InlineKeyboardMarkup{Keyboard: keyboard},
	})
	if err != nil {
		return err
	}

	err = s.telegramApi.DeleteMessage(telegramapi.DeleteMessageRequest{
		ChatId:    s.chatId,
		MessageId: s.messageId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *TelegramService) Response(params models.TelegramUpdate) (err error) {
	if params.Message != nil && params.Message.From.Username != "looken15" {
		return
	}

	if params.CallbackQuery != nil && params.CallbackQuery.From.Username != "looken15" {
		return
	}

	if params.Message != nil {
		s.fillParams(params.Message.From.ID, params.Message.MessageID, params.Message.Chat.ID, params.Message.From.Username)
	}
	if params.CallbackQuery != nil {
		s.fillParams(params.CallbackQuery.From.ID, params.CallbackQuery.Message.MessageID, params.CallbackQuery.Message.Chat.ID, params.CallbackQuery.From.Username)
	}
	if params.CallbackQuery != nil {
		err = s.fillCallbackData(params.CallbackQuery.Data)
		if err != nil {
			return err
		}

		_ = s.telegramApi.AnswerCallbackQuery(telegramapi.AnswerCallbackQueryRequest{
			CallbackQueryId: params.CallbackQuery.Id,
		})
	}

	isUserSub, err := s.telegramApi.CheckIfUserSub(s.userId)
	if err != nil {
		return
	}

	if !isUserSub {
		return s.sendUnsubMessage()
	}

	if params.CallbackQuery != nil && lo.Contains(_positionsArray, s.callbackData.Position) {
		return s.sendPositionsMessage()
	}

	if params.CallbackQuery != nil && lo.Contains(_tacticsArray, s.callbackData.Tactic) {
		return s.sendTacticMessage()
	}

	if s.callbackData.NextCommand == _tryAgainMessage || (params.Message != nil && params.Message.Text == _startMessage) || s.callbackData.NextCommand == _startMessage {
		return s.sendStartMessage()
	}

	if s.callbackData.NextCommand == _toPositionsMessage || (params.Message != nil && params.Message.Text == _toPositionsMessage) {
		return s.sendPositionsPickMessage()
	}

	if s.callbackData.NextCommand == _toTacticsMessage || (params.Message != nil && params.Message.Text == _toTacticsMessage) {
		return s.sendTacticsPickMessage()
	}

	return
}
