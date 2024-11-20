package transcriber

import (
	"context"
	transcriptionClient "github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/transcriber"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/models/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/service"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type serv struct {
	client transcriptionClient.TranscriptionServiceClient
}

func New(client transcriptionClient.TranscriptionServiceClient) service.TranscriberService {
	return &serv{
		client: client,
	}
}

func (t *serv) TranscribeByPath(ctx context.Context, filePath string, lang *string) (string, error) {
	res, err := t.client.TranscribeByPath(
		ctx, &transcriptionClient.TranscribeByPathRequest{
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

func (t *serv) TranscribeByBinary(ctx context.Context, audioData []byte, lang *string) (string, error) {
	res, err := t.client.TranscribeByBinary(
		ctx, &transcriptionClient.TranscribeByBinaryRequest{
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

func (t *serv) AvailableLanguages(ctx context.Context) ([]string, error) {
	res, err := t.client.GetAvailableLanguages(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get available languages")
	}

	return res.Languages, nil
}
