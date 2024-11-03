package storage

import (
	"context"
)

type Storage interface {
	SaveUnprocessed(ctx context.Context, event *Event) error
	SetProcessed(ctx context.Context, event *Event) error
	DeleteProcessed(ctx context.Context, event *Event) error
	FetchUnprocessed(ctx context.Context, limit int) ([]Event, error)
}

type Event struct {
	ID        uint
	Username  string
	FilePath  string
	Processed bool
}
