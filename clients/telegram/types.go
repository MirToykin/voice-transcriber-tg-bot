package telegram

type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type File struct {
	FilePath string `json:"file_path"`
}

type Voice struct {
	FileID        string
	DurationSec   int
	MimeType      string
	FileSizeBytes int
}

type IncomingMessage struct {
	Text  *string
	Voice *Voice `json:"voice"`
	From  From   `json:"from"`
	Chat  Chat   `json:"chat"`
}
