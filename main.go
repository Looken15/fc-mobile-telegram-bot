package main

import (
	"context"
	"fc-mobile-telegram-bot/config"
	"fc-mobile-telegram-bot/db"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

func main() {
	decimal.DivisionPrecision = 2
	decimal.MarshalJSONWithoutQuotes = true

	settings := config.Get()
	db.RunMigrate(settings)

	mainCtx, cancelMainCtx := context.WithCancel(context.Background())
	defer cancelMainCtx()

	app := NewApp(mainCtx, settings)
	app.Init()
	app.Start()

	stop := make(chan os.Signal, 1)
	<-stop

	if err := app.Stop(context.WithTimeout); err != nil {
		log.Fatal(err)
		return
	}
}
