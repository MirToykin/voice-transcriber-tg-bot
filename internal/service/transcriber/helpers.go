package transcriber

import (
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/transcriber"
	"net/http"
)

func checkIfTranscriptionNeedRetry(res *transcription.TranscriptionResponse) bool {
	return !res.GetStatus() && res.GetErrorCode() != http.StatusNotFound
}
