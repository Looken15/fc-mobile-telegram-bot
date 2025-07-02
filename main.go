package main

import (
	"context"
	"github.com/shopspring/decimal"
	"log"
	"os"
)

func main() {
	decimal.DivisionPrecision = 2
	decimal.MarshalJSONWithoutQuotes = true

	mainCtx, cancelMainCtx := context.WithCancel(context.Background())
	defer cancelMainCtx()

	app := NewApp(mainCtx)
	app.Init()
	app.Start()

	stop := make(chan os.Signal, 1)
	<-stop

	if err := app.Stop(context.WithTimeout); err != nil {
		log.Fatal(err)
		return
	}
}
