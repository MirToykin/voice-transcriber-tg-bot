package storage

import (
	"encoding/json"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/pkg/errors"
)

func FromStorageToBaseEvent(event *Event) (*events.Event, error) {
	var meta interface{}
	err := json.Unmarshal([]byte(event.Meta), meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert storage event to base event")
	}
	return &events.Event{
		Type: event.Type,
		AudioFile: events.AudioFile{
			Path:      event.FilePath,
			SizeBytes: event.FileSize,
		},
		Text: event.Text,
		Meta: meta,
	}, nil
}

func FromBaseToStorageEvent(event *events.Event) (*Event, error) {
	data, err := json.Marshal(event.Meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert base event to storage event")
	}
	return &Event{
		Type:     event.Type,
		FilePath: event.AudioFile.Path,
		FileSize: event.AudioFile.SizeBytes,
		Text:     event.Text,
		Meta:     string(data),
	}, nil
}
