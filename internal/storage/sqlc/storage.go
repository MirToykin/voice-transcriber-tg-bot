package sqlc

import (
	"database/sql"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/storage/sqlc/db"
	"github.com/gomscourse/common/pkg/closer"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	qrs *db.Queries
}

func New(storagePath string) (*Storage, error) {
	database, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	closer.Add(database.Close)

	queries := db.New(database)

	s := &Storage{qrs: queries}

	return s, nil
}
