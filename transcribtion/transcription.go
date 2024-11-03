package transcribtion

import "context"

type Transcriber interface {
	TranscribeByPath(ctx context.Context, filePath string) (string, error)
	TranscribeByBinary(ctx context.Context, audioData []byte) (string, error)
}
