package telegram

import (
	"context"
	tgClient "github.com/MirToykin/voice-transcriber-tg-bot/clients/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion"
	"github.com/pkg/errors"
	"log"
)

type Processor struct {
	tgClient    tgClient.Client
	offset      int
	transcriber transcribtion.Transcriber
}

type Meta struct {
	ChatID int
	User   tgClient.From
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client tgClient.Client, transcriber transcribtion.Transcriber) Processor {
	return Processor{
		tgClient:    client,
		transcriber: transcriber,
	}
}

func (p *Processor) Fetch(ctx context.Context, limit int) ([]events.Event, error) {
	updates, err := p.tgClient.Updates(ctx, p.offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tg updates")
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for i, u := range updates {
		ev, err := updateToEvent(ctx, u, p.tgClient)

		if err != nil {
			log.Printf("ERROR: failed to convert update %d to event: %s\n", u.ID, err)
			p.offset = updates[i].ID
			break
		}

		res = append(res, *ev)
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(ctx context.Context, e *events.Event) error {
	switch e.Type {
	case events.TextMessage:
		return p.processTextMessage(ctx, e)
	case events.VoiceMessage:
		return p.processVoiceMessage(ctx, e)
	default:
		return errors.Wrap(ErrUnknownEventType, "can't process message")
	}
}

func (p *Processor) processTextMessage(ctx context.Context, e *events.Event) error {
	meta, err := toTgProcessorMeta(e)
	if err != nil {
		return errors.Wrap(err, "failed to process text message")
	}

	err = p.doTextCmd(ctx, e.Text, meta.ChatID, &meta.User)
	if err != nil {
		return errors.Wrap(err, "failed to perform text cmd")
	}

	return nil
}

func (p *Processor) processVoiceMessage(ctx context.Context, e *events.Event) error {
	meta, err := toTgProcessorMeta(e)
	if err != nil {
		return errors.Wrap(err, "failed to process voice message")
	}

	var langCode *string
	if meta.User.LanguageCode != "" {
		langCode = &meta.User.LanguageCode
	}

	return p.sendTranscription(ctx, meta.ChatID, e.AudioFilePath, langCode)
}
