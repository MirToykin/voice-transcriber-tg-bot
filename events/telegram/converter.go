package telegram

import (
	"context"
	"github.com/pkg/errors"
	"path"
	tgClient "voice_transcriber_bot/clients/telegram"
	"voice_transcriber_bot/events"
)

func updateToEvent(ctx context.Context, upd tgClient.Update, client tgClient.Client) (*events.Event, error) {
	updType := fetchType(upd)
	filePath, err := fetchFilePath(ctx, upd, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch file path")
	}

	res := &events.Event{
		Type:          updType,
		AudioFilePath: filePath,
	}

	if updType == events.TextMessage || updType == events.VoiceMessage {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
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
		if upd.Message.Text != nil {
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

	return path.Join(client.FilesPath(), file.FilePath), nil
}
