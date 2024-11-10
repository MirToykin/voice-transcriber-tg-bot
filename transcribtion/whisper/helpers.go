package whisper

import (
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion/whisper/generated"
	"net/http"
)

func checkIfTranscriptionNeedRetry(res *generated.TranscriptionResponse) bool {
	return !res.GetStatus() && res.GetErrorCode() != http.StatusNotFound
}
