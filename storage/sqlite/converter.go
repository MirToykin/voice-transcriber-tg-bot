package sqlite

import "voice_transcriber_bot/storage"

func toStorageEvent(e Event) storage.Event {
	return storage.Event{
		ID:        e.ID,
		Username:  e.Username,
		FilePath:  e.FilePath,
		Processed: e.Processed,
	}
}
