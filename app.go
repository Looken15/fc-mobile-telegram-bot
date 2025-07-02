package main

import (
	"context"
	"fc-mobile-telegram-bot/api"
	"fmt"
	"net/http"
	"time"
)

type App struct {
	server  *http.Server
	mainCtx context.Context
}

func NewApp(mainCtx context.Context) *App {
	return &App{
		mainCtx: mainCtx,
	}
}

func (app *App) Start() {
	go func() {
		if err := app.server.ListenAndServe(); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()
}

func (app *App) Stop(getContext func(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)) error {
	serverCtx, cancelServerCtx := getContext(app.mainCtx, time.Second*15)
	defer cancelServerCtx()

	err := app.server.Shutdown(serverCtx)
	if err != nil {
		fmt.Println("Error shutting down server:", err)
		return err
	}

	return nil
}

func (app *App) Init() {
	app.server = api.NewServer(
		app.mainCtx)
}
