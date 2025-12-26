package db

import (
	"fc-mobile-telegram-bot/config"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"path/filepath"
)

func RunMigrate(settings *config.Settings) {
	fmt.Println(settings.DbUrl)

	db, err := sqlx.Open("pgx", settings.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(cwd, "migrations")

	fmt.Println(dir)

	if err := goose.Up(db.DB, dir); err != nil {
		log.Fatal(err)
	}
}
