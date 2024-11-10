package whisper

import (
	"context"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion/whisper/generated"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
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
		ctx, &generated.TranscribeByPathRequest{
			FilePath: filePath,
			Lang:     lang,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to transcribe by path")
	}

	if !res.GetStatus() {
		return "", events.NewProcessingError(res.GetErrorDescription(), checkIfTranscriptionNeedRetry(res), nil)
	}

	return res.GetText(), nil
}

func (t *Transcriber) TranscribeByBinary(ctx context.Context, audioData []byte, lang *string) (string, error) {
	res, err := t.client.TranscribeByBinary(
		ctx, &generated.TranscribeByBinaryRequest{
			AudioData: audioData,
			Lang:      lang,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to transcribe by binary")
	}

	if !res.GetStatus() {
		return "", events.NewProcessingError(res.GetErrorDescription(), checkIfTranscriptionNeedRetry(res), nil)
	}

	return res.GetText(), nil
}

func (t *Transcriber) AvailableLanguages(ctx context.Context) ([]string, error) {
	res, err := t.client.GetAvailableLanguages(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get available languages")
	}

	return res.Languages, nil
}
