package telegram

import (
	"context"
	tgClient "voice_transcriber_bot/clients/telegram"
	"voice_transcriber_bot/events"
	"voice_transcriber_bot/transcribtion"
)

type Processor struct {
	tgClient    tgClient.Client
	offset      int
	transcriber transcribtion.Transcriber
}

func New(client tgClient.Client, transcriber transcribtion.Transcriber) Processor {
	return Processor{
		tgClient:    client,
		transcriber: transcriber,
	}
}

func (p *Processor) Fetch(ctx context.Context, limit int) ([]events.Event, error) {
	return nil, nil
}

func (p *Processor) Process(ctx context.Context, e *events.Event) error {
	return nil
}
