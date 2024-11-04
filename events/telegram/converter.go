package telegram

import (
	"context"
	tgClient "github.com/MirToykin/voice-transcriber-tg-bot/clients/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/pkg/errors"
)

func updateToEvent(ctx context.Context, upd tgClient.Update, client tgClient.Client) (*events.Event, error) {
	updType := fetchType(upd)
	filePath, err := fetchFilePath(ctx, upd, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch file path")
	}

	res := &events.Event{
		ID:            upd.ID,
		Type:          updType,
		AudioFilePath: filePath,
		Text:          fetchText(upd),
	}

	if updType == events.TextMessage || updType == events.VoiceMessage {
		res.Meta = Meta{
			ChatID:    upd.Message.Chat.ID,
			User:      upd.Message.From,
			MessageID: upd.Message.ID,
		}
	}

	return res, nil
}

func toTgProcessorMeta(event *events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, errors.Wrap(ErrUnknownMetaType, "can't get meta")
	}

	return res, nil
}

func fetchType(upd tgClient.Update) events.Type {
	if upd.Message != nil {
		if upd.Message.Text != nil && upd.Message.Voice == nil {
			return events.TextMessage
		} else if upd.Message.Voice != nil {
			return events.VoiceMessage
		}
	}

	return events.Unknown
}

func fetchFilePath(ctx context.Context, upd tgClient.Update, client tgClient.Client) (string, error) {
	if upd.Message == nil || upd.Message.Voice == nil {
		return "", nil
	}

	file, err := client.File(ctx, upd.Message.Voice.FileID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get file info")
	}

	return client.FileFullPath(file.FilePath), nil
}

func fetchText(upd tgClient.Update) string {
	if upd.Message == nil || upd.Message.Text == nil {
		return ""
	}

	return *upd.Message.Text
}
