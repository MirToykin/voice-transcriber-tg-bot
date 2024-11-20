package service

import (
	"context"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/models/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/models/users"
)

type Fetcher interface {
	Fetch(ctx context.Context, limit int) ([]*events.Event, error)
}

type Processor interface {
	Process(ctx context.Context, e *events.Event) error
}

type EventService interface {
	Fetcher
	Processor
}

type TranscriberService interface {
	TranscribeByPath(ctx context.Context, filePath string, lang *string) (string, error)
	TranscribeByBinary(ctx context.Context, audioData []byte, lang *string) (string, error)
	AvailableLanguages(ctx context.Context) ([]string, error)
}

type UserService interface {
	CreateUser(ctx context.Context, params users.UserCreateParams)
}
