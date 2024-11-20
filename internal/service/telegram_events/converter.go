package telegram_events

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/models/events"
	"github.com/pkg/errors"
)

func updateToEvent(ctx context.Context, upd telegram.Update, client telegram.Client) (*events.Event, error) {
	updType := fetchType(upd)
	filePath, err := fetchFilePath(ctx, upd, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch file path")
	}

	res := &events.Event{
		ExternalID: upd.ID,
		Type:       updType,
		AudioFile: events.AudioFile{
			Path:      filePath,
			SizeBytes: fetchFileSize(upd),
		},
		Text: fetchText(upd),
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
		data, err := json.Marshal(event.Meta)
		if err != nil {
			return Meta{}, errors.Wrap(
				ErrUnknownMetaType,
				fmt.Sprintf("unable to get meta (second try): %s", err.Error()),
			)
		}

		err = json.Unmarshal(data, &res)
		if err != nil {
			return Meta{}, errors.Wrap(
				ErrUnknownMetaType,
				fmt.Sprintf("unable to get meta (second try): %s", err.Error()),
			)
		}
	}

	if res.User.LanguageCode == "" {
		res.User.LanguageCode = defaultLanguage
	}

	return res, nil
}

func fetchType(upd telegram.Update) events.EventType {
	if upd.Message != nil {
		if upd.Message.Text != nil && upd.Message.Voice == nil {
			if upd.Message.Chat.Type == events.GroupChat {
				return events.GroupTextMessage
			}
			return events.TextMessage
		} else if upd.Message.Voice != nil {
			return events.VoiceMessage
		}
	}

	return events.Unknown
}

func fetchFilePath(ctx context.Context, upd telegram.Update, client telegram.Client) (string, error) {
	if upd.Message == nil || upd.Message.Voice == nil {
		return "", nil
	}

	file, err := client.File(ctx, upd.Message.Voice.FileID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get file info")
	}

	return client.FileFullPath(file.FilePath), nil
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil || upd.Message.Text == nil {
		return ""
	}

	return *upd.Message.Text
}

func fetchFileSize(upd telegram.Update) int {
	if upd.Message == nil || upd.Message.Voice == nil {
		return 0
	}

	return upd.Message.Voice.FileSizeBytes
}
