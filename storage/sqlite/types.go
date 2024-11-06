package sqlite

import "github.com/MirToykin/voice-transcriber-tg-bot/events"

type Event struct {
	ID        uint        `db:"id"`
	Type      events.Type `db:"type"`
	FilePath  string      `db:"file_path"`
	FileSize  int         `db:"file_size"`
	Text      string      `db:"text"`
	Meta      string      `db:"meta"`
	Processed bool        `db:"processed"`
}
