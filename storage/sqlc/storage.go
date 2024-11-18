package sqlc

import (
	"database/sql"
	"github.com/MirToykin/voice-transcriber-tg-bot/storage/sqlc/db"
)

type Storage struct {
	qrs *db.Queries
}

func New(storagePath string) (*Storage, error) {
	database, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	queries := db.New(database)

	s := &Storage{qrs: queries}

	return s, nil
}
