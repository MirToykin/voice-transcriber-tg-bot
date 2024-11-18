package sqlc

import (
	"encoding/json"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/storage/sqlc/db"
	"github.com/pkg/errors"
)

func fromEventToBase(event *db.FetchUnprocessedEventsRow) (*events.Event, error) {
	baseEvent := &events.Event{
		ID:   uint(event.ID),
		Type: events.EventType(event.Type),
		AudioFile: events.AudioFile{
			Path:      event.FilePath.String,
			SizeBytes: int(event.FileSize.Int64),
		},
		Text: event.Text.String,
	}

	var meta interface{}
	err := json.Unmarshal([]byte(event.Meta.String), &meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal event meta")
	}

	baseEvent.Meta = meta

	return baseEvent, nil
}
