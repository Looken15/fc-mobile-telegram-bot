package handlers

import (
	"encoding/json"
	"fc-mobile-telegram-bot/models"
	"fc-mobile-telegram-bot/service/telegramservice"
	"fc-mobile-telegram-bot/utils"
	"log"
	"net/http"
)

func UpdateHandler(telegramService telegramservice.TelegramService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := models.TelegramUpdate{}
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_ = utils.PrintJSON(params)

		err = telegramService.Response(params)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
