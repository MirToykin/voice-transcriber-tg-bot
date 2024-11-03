package whisper

import (
	"context"
	"github.com/pkg/errors"
	"voice_transcriber_bot/transcribtion/whisper/generated"
)

type Transcriber struct {
	client generated.TranscriptionServiceClient
}

func New(client generated.TranscriptionServiceClient) *Transcriber {
	return &Transcriber{
		client: client,
	}
}

func (t *Transcriber) TranscribeByPath(ctx context.Context, filePath string) (string, error) {
	res, err := t.client.TranscribeByPath(
		ctx, &generated.TranscribePathRequest{
			FilePath: filePath,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to transcribe by path")
	}

	return res.GetText(), nil
}

func (t *Transcriber) TranscribeByBinary(ctx context.Context, audioData []byte) (string, error) {
	res, err := t.client.TranscribeByBinary(
		ctx, &generated.TranscribeBinaryRequest{
			AudioData: audioData,
		},
	)

	if err != nil {
		return "", errors.Wrap(err, "failed to transcribe by binary")
	}

	return res.GetText(), nil
}
