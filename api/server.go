package api

import (
	"context"
	"fc-mobile-telegram-bot/api/handlers"
	"fc-mobile-telegram-bot/config"
	"fc-mobile-telegram-bot/service/telegramservice"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func NewServer(ctx context.Context, settings *config.Settings, service telegramservice.TelegramService) *http.Server {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/ping", handlers.PingHandler()).Methods(http.MethodGet)
	router.HandleFunc("/update", handlers.UpdateHandler(service)).Methods(http.MethodPost)

	return &http.Server{
		Addr:        fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context { return ctx },
		Handler:     router,
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
