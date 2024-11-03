package config

import (
	"log"
	"os"
)

type Config struct {
	TgBotToken      string
	DbDSN           string
	TranscriberHost string
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

	trbHost := os.Getenv(transcriberHost)
	if trbHost == "" {
		log.Fatal("transcriber host isn't set")
	}

	return Config{
		TgBotToken:      token,
		DbDSN:           dsn,
		TranscriberHost: trbHost,
	}
}
