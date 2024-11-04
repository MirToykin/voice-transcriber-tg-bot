package storage

import (
	"context"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
)

type Storage interface {
	SaveUnprocessed(ctx context.Context, event *events.Event) error
	SetProcessed(ctx context.Context, eventId uint) error
	DeleteProcessed(ctx context.Context) error
	FetchUnprocessed(ctx context.Context, limit int) ([]events.Event, error)
}

type Event struct {
	ID        uint
	Type      events.Type
	FilePath  string
	Text      string
	Meta      string
	Processed bool
}
