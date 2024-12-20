package sqlite

import (
	"context"
	"database/sql"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/lib/e"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"log"
)

type Storage struct {
	db *sqlx.DB
}

func New(ctx context.Context, storagePath string) (*Storage, error) {
	db, err := sqlx.Open("sqlite3", storagePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "failed to connect to database: %s", storagePath)
	}

	s := &Storage{db: db}

	err = s.Init(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init database")
	}

	return s, nil
}

func (s *Storage) SaveUnprocessed(ctx context.Context, event *events.Event) (err error) {
	defer func() { err = e.WrapIfErr("failed to save unprocessed event", err) }()

	q := "INSERT INTO events (type, file_path, file_size, text, meta, processed) values (?, ?, ?, ?, ?, 0)"

	_, err = s.db.ExecContext(
		ctx,
		q,
		event.Type,
		event.AudioFile.Path,
		event.AudioFile.SizeBytes,
		event.Text,
		event.StringMeta(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SetProcessed(ctx context.Context, eventId uint) error {
	q := "UPDATE events SET processed = 1 WHERE id = ?"

	_, err := s.db.ExecContext(ctx, q, eventId)
	if err != nil {
		return errors.Wrap(err, "failed to set event processed")
	}

	return nil
}

func (s *Storage) DeleteProcessed(ctx context.Context) error {
	q := "DELETE FROM events WHERE processed = 1"

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return errors.Wrap(err, "failed to save unprocessed event")
	}

	return nil
}

func (s *Storage) FetchUnprocessed(ctx context.Context, limit int) ([]*events.Event, error) {
	q := "SELECT id, type, file_path, file_size, text, meta FROM events WHERE processed = 0 LIMIT ?"

	var eventsList []Event

	err := s.db.SelectContext(ctx, &eventsList, q, limit)

	if errors.Is(err, sql.ErrNoRows) {
		return []*events.Event{}, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch processed")
	}

	storageEvents := make([]*events.Event, 0, len(eventsList))
	for _, ev := range eventsList {
		baseEvent, err := fromEventToBase(&ev)
		if err != nil {
			log.Println("ERROR: failed to convert database event to event")
			continue
		}
		storageEvents = append(storageEvents, baseEvent)
	}

	return storageEvents, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `
CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	type INTEGER,
	file_path TEXT,
	file_size integer,
	text TEXT,
	meta TEXT,
	processed BOOLEAN NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_events_processed ON events (processed);`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return errors.Wrap(err, "failed to initiate database")
	}

	return nil
}
