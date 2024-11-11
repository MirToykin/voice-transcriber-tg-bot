package events

import (
	"context"
	"encoding/json"
	"fmt"
)

type Fetcher interface {
	Fetch(ctx context.Context, limit int) ([]*Event, error)
}

type Processor interface {
	Process(ctx context.Context, e *Event) error
}

type AudioFile struct {
	Path      string
	SizeBytes int
}

type Event struct {
	ID         uint
	ExternalID any
	Type       EventType
	AudioFile  AudioFile
	Text       string
	Meta       interface{}
}

func (e *Event) String() string {
	return fmt.Sprintf(
		"ID: %v, ExternalID: %v, Type: %d, AudioPath: %s, AudioSize: %d, Text: %s",
		e.ID,
		e.ExternalID,
		e.Type,
		e.AudioFile.Path,
		e.AudioFile.SizeBytes,
		e.Text,
	)
}

func (e *Event) StringMeta() string {
	data, err := json.Marshal(&e.Meta)
	if err != nil {
		fmt.Printf("failed to stringify event meta (event: %s)\n", e)
	}

	return string(data)
}
