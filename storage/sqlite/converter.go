package sqlite

import (
	"encoding/json"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/pkg/errors"
)

func fromBaseToEvent(event *events.Event) (*Event, error) {
	stEvent := &Event{
		Type:     event.Type,
		FilePath: event.AudioFilePath,
		Text:     event.Text,
	}

	stMeta, err := json.Marshal(event.Meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal event meta")
	}

	stEvent.Meta = string(stMeta)

	return stEvent, nil
}

func fromEventToBase(event *Event) (*events.Event, error) {
	baseEvent := &events.Event{
		Type:          event.Type,
		AudioFilePath: event.FilePath,
		Text:          event.Text,
	}

	var meta interface{}
	err := json.Unmarshal([]byte(event.Meta), &meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal event meta")
	}

	baseEvent.Meta = meta

	return baseEvent, nil
}
