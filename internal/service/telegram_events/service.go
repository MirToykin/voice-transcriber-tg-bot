package telegram_events

import (
	"context"
	"fmt"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/models/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/service"
	"github.com/MirToykin/voice-transcriber-tg-bot/pkg/lib/e"
	"github.com/pkg/errors"
	"log"
)

const (
	allowedSizeMB  = 10
	showInfoSizeMB = 3
)

type serv struct {
	tgClient    telegram.Client
	offset      int
	transcriber service.TranscriberService
}

type Meta struct {
	ChatID    int
	User      telegram.From
	MessageID int
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(client telegram.Client, transcriber service.TranscriberService) service.EventService {
	return &serv{
		tgClient:    client,
		transcriber: transcriber,
	}
}

func (p *serv) Fetch(ctx context.Context, limit int) ([]*events.Event, error) {
	updates, err := p.tgClient.Updates(ctx, p.offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tg updates")
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]*events.Event, 0, len(updates))

	for i, u := range updates {
		ev, err := updateToEvent(ctx, u, p.tgClient)

		if err != nil {
			log.Printf("ERROR: failed to convert update %d to event: %s\n", u.ID, err)
			p.offset = updates[i].ID
			break
		}

		res = append(res, ev)
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *serv) Process(ctx context.Context, e *events.Event) error {
	switch e.Type {
	case events.TextMessage:
		return p.processTextMessage(ctx, e)
	case events.VoiceMessage:
		return p.processVoiceMessage(ctx, e)
	case events.GroupTextMessage:
		return nil
	default:
		return &events.ProcessingError{
			Message:   "can't process message",
			Cause:     ErrUnknownEventType,
			NeedRetry: false,
		}
	}
}

func (p *serv) processTextMessage(ctx context.Context, e *events.Event) error {
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

func (p *serv) processVoiceMessage(ctx context.Context, evt *events.Event) (err error) {
	defer func() { err = e.WrapIfErr("failed to process voice message", err) }()
	meta, err := toTgProcessorMeta(evt)
	if err != nil {
		return err
	}

	return p.sendMsgDependingOnFileSize(ctx, evt, meta)
}

func (p *serv) sendMsgDependingOnFileSize(ctx context.Context, evt *events.Event, meta Meta) error {
	fileSize := evt.AudioFile.SizeBytes
	if fileSize > getBytesSize(allowedSizeMB) {
		return p.sendFileSizeExceeded(ctx, meta)
	} else if fileSize >= getBytesSize(showInfoSizeMB) {
		return p.sendWithInfo(ctx, evt, meta)
	} else {
		return p.sendTranscription(ctx, meta, evt.AudioFile.Path, fetchLanguageCode(meta.User.LanguageCode))
	}
}

func (p *serv) sendFileSizeExceeded(ctx context.Context, meta Meta) error {
	return p.tgClient.SendReplyMessage(
		ctx,
		meta.ChatID,
		fmt.Sprintf("Разрешенный размер файла до %d МБ", allowedSizeMB),
		meta.MessageID,
	)
}

func (p *serv) sendWithInfo(ctx context.Context, evt *events.Event, meta Meta) error {
	err := p.tgClient.SendReplyMessage(
		ctx,
		meta.ChatID,
		"Идет распознавание файла, процесс может занять некоторое время.",
		meta.MessageID,
	)

	if err != nil {
		log.Printf("failed to send info message: %s\n", err.Error())
	}

	return p.sendTranscription(ctx, meta, evt.AudioFile.Path, fetchLanguageCode(meta.User.LanguageCode))
}

func fetchLanguageCode(userLanguageCode string) *string {
	var langCode *string

	if userLanguageCode != "" {
		langCode = &userLanguageCode
	}

	return langCode
}

func getBytesSize(mbSize int) int {
	return mbSize * 1024 * 1024
}
