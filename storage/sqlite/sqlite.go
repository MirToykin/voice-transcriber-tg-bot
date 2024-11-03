package sqlite

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"voice_transcriber_bot/storage"
)

type Storage struct {
	db *sqlx.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sqlx.Open("sqlite3", storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUnprocessed(ctx context.Context, event *storage.Event) error {
	q := "INSERT INTO events (username, file_path, processed) values (?, ?, 0)"

	_, err := s.db.ExecContext(ctx, q, event.Username, event.FilePath)
	if err != nil {
		return errors.Wrap(err, "failed to save unprocessed event")
	}

	return nil
}

func (s *Storage) SetProcessed(ctx context.Context, event *storage.Event) error {
	q := "UPDATE events SET processed = 1 WHERE id = ?"

	_, err := s.db.ExecContext(ctx, q, event.ID)
	if err != nil {
		return errors.Wrap(err, "failed to set event processed")
	}

	return nil
}

func (s *Storage) DeleteProcessed(ctx context.Context, event *storage.Event) error {
	q := "DELETE FROM events WHERE processed = 1"

	_, err := s.db.ExecContext(ctx, q, event.ID)
	if err != nil {
		return errors.Wrap(err, "failed to save unprocessed event")
	}

	return nil
}

func (s *Storage) FetchUnprocessed(ctx context.Context, limit int) ([]storage.Event, error) {
	q := "SELECT id, username, file_path FROM unprocessed_events WHERE processed = 0 LIMIT ?"

	var events []Event

	err := s.db.SelectContext(ctx, &events, q, limit)

	if errors.Is(err, sql.ErrNoRows) {
		return []storage.Event{}, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch processed")
	}

	stEvents := make([]storage.Event, len(events))
	for i, ev := range events {
		stEvents[i] = toStorageEvent(ev)
	}

	return stEvents, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `
CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	file_path TEXT NOT NULL,
	processed BOOLEAN NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_events_processed ON events (processed);`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return errors.Wrap(err, "failed to initiate database")
	}

	return nil
}
