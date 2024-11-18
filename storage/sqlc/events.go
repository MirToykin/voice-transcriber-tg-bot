package sqlc

import (
	"context"
	"database/sql"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/storage/sqlc/db"
	"github.com/pkg/errors"
	"log"
)

func (s *Storage) SaveUnprocessed(ctx context.Context, event *events.Event) error {
	err := s.qrs.SaveUnprocessedEvent(
		ctx, db.SaveUnprocessedEventParams{
			Type:     int64(event.Type),
			FilePath: stringToSQLNullString(event.AudioFile.Path),
			FileSize: intToSQLNullInt(event.AudioFile.SizeBytes),
			Text:     stringToSQLNullString(event.Text),
			Meta:     stringToSQLNullString(event.StringMeta()),
		},
	)

	if err != nil {
		return errors.Wrap(err, "failed to save unprocessed event")
	}

	return nil
}

func (s *Storage) SetProcessed(ctx context.Context, eventId uint) error {
	err := s.qrs.SetEventProcessed(ctx, int64(eventId))
	if err != nil {
		return errors.Wrap(err, "failed to set event processed")
	}

	return nil
}

func (s *Storage) DeleteProcessed(ctx context.Context) error {
	err := s.qrs.DeleteProcessedEvents(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to save unprocessed event")
	}

	return nil
}

func (s *Storage) FetchUnprocessed(ctx context.Context, limit int) ([]*events.Event, error) {
	eventsList, err := s.qrs.FetchUnprocessedEvents(ctx, int64(limit))

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
