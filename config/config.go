package config

import (
	"log"
	"os"
)

type Config struct {
	TgBotToken string
	DbDSN      string
}

func MustLoad() Config {
	token := os.Getenv(tgBotToken)
	if token == "" {
		log.Fatal("Telegram bot token isn't set")
	}

	dsn := os.Getenv(dbDSN)
	if dsn == "" {
		log.Fatal("database connection isn't set")
	}

	return Config{
		TgBotToken: token,
		DbDSN:      dsn,
	}
}
