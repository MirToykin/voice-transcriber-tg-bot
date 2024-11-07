package whisper

import (
	"context"
	"fmt"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion/whisper/generated"
	"github.com/pkg/errors"
	"strings"
)

type Transcriber struct {
	client generated.TranscriptionServiceClient
}

func New(client generated.TranscriptionServiceClient) *Transcriber {
	return &Transcriber{
		client: client,
	}
}

func (t *Transcriber) TranscribeByPath(ctx context.Context, filePath string, lang *string) (string, error) {
	res, err := t.client.TranscribeByPath(
		ctx, &generated.TranscribePathRequest{
			FilePath: filePath,
			Lang:     lang,
		},
	)

	if err != nil {
		//TODO: сделать более надежный способ проверки
		if strings.Contains(err.Error(), fmt.Sprintf("Not Found for url: %s", filePath)) {
			err = &events.ProcessingError{
				Message:   "file not found",
				Cause:     err,
				NeedRetry: false,
			}
		}
		return "", errors.Wrap(err, "failed to transcribe by path")
	}

	return res.GetText(), nil
}

func (t *Transcriber) TranscribeByBinary(ctx context.Context, audioData []byte, lang *string) (string, error) {
	res, err := t.client.TranscribeByBinary(
		ctx, &generated.TranscribeBinaryRequest{
			AudioData: audioData,
			Lang:      lang,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to transcribe by binary")
	}

	return res.GetText(), nil
}
