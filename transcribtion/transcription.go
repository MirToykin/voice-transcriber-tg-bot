package transcribtion

import "context"

type Transcriber interface {
	TranscribeByPath(ctx context.Context, filePath string, lang *string) (string, error)
	TranscribeByBinary(ctx context.Context, audioData []byte, lang *string) (string, error)
}
