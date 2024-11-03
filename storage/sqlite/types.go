package sqlite

import "voice_transcriber_bot/events"

type Event struct {
	ID        uint        `db:"id"`
	Type      events.Type `db:"type"`
	FilePath  string      `db:"file_path"`
	Text      string      `db:"text"`
	Meta      string      `db:"meta"`
	Processed bool        `db:"processed"`
}
