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
	_imagePathJPG  = "./images/%s.jpg"
	_imagePathPNG  = "./images/%s.png"

	_lastUpdateDate = "22 декабря, 2025"

	_sendPhotoCaption = "<b>ТОП-10 %s в FC Mobile</b>\n\nПоследнее обновление:\n%s\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"

	_subscribeNeededCaption = "Чтобы использовать бота, необходимо подписаться на каналы <a href=\"https://t.me/+mf4AwsUOHlBiNDky\"> KARAMA | FC MOBILE 26 | FIFA MOBILE </a> и <a href=\"https://t.me/+rkUjX8CQwYcwMjQy\"> BASEMENT ATHLETIC | FC MOBILE </a> и нажать кнопку «Проверить подписку»"

	_helloCaption = "<b>Приветствую, @%s.</b>\n\nВ этом боте вы найдете ТОП-10 игроков на каждую позицию\n\n<a href=\"http://t.me/KaramaFC\">KARAMA | FC MOBILE 26</a>"
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

func (s *TelegramService) sendPositionsMessage() error {
	position := s.callbackData.Position

	keyboard := make([][]telegramapi.InlineKeyboardButton, 0)
	keyboardLine := make([]telegramapi.InlineKeyboardButton, 0)

	newCallbackData := utils.CallbackData{
		NextCommand: _backMessage,
		MessageId:   s.messageId,
	}
	keyboardLine = append(keyboardLine, telegramapi.InlineKeyboardButton{
		Text:         "Назад",
		CallbackData: utils.EncodeCallbackData(newCallbackData),
	})
	keyboard = append(keyboard, keyboardLine)

	_, err := s.telegramApi.SendPhoto(telegramapi.SendPhotoRequest{
		ChatId:               s.chatId,
		Caption:              fmt.Sprintf(_sendPhotoCaption, _positionsWordMap[position], _lastUpdateDate),
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

func (s *TelegramService) Response(params models.TelegramUpdate) (err error) {
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
		err = s.sendUnsubMessage()
		if err != nil {
			return err
		}
	}

	if params.CallbackQuery != nil && lo.Contains(_positionsArray, s.callbackData.Position) {
		err = s.sendPositionsMessage()
		if err != nil {
			return err
		}
	}

	if s.callbackData.NextCommand == _tryAgainMessage || s.callbackData.NextCommand == _backMessage || (params.Message != nil && params.Message.Text == _startMessage) {
		err = s.sendStartMessage()
		if err != nil {
			return err
		}
	}

	return
}
