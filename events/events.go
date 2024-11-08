package events

import (
	"context"
	"fmt"
)

type ProcessingError struct {
	Message   string
	NeedRetry bool
	Cause     error
}

func (e *ProcessingError) Error() string {
	needRetry := "false"
	if e.NeedRetry {
		needRetry = "true"
	}

	if e.Cause != nil {
		return fmt.Sprintf("%s, need retry: %s, caused by: %v", e.Message, needRetry, e.Cause)
	}
	return fmt.Sprintf("%s, need retry: %s", e.Message, needRetry)
}

func (e *ProcessingError) Unwrap() error {
	return e.Cause
}

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
	ID        any
	Type      EventType
	AudioFile AudioFile
	Text      string
	Meta      interface{}
}

func (e *Event) String() string {
	return fmt.Sprintf(
		"ID: %v, Type: %d, AudioPath: %s, AudioSize: %d, Text: %s",
		e.ID,
		e.Type,
		e.AudioFile.Path,
		e.AudioFile.SizeBytes,
		e.Text,
	)
}
