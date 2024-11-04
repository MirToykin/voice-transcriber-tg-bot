package telegram

type BaseResponse struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type UpdatesResponse struct {
	BaseResponse
	Result []Update `json:"result"`
}

type FileResponse struct {
	BaseResponse
	Result File
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type From struct {
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID int `json:"id"`
}

type File struct {
	FilePath string `json:"file_path"`
}

type Voice struct {
	FileID        string `json:"file_id"`
	DurationSec   int    `json:"duration"`
	MimeType      string `json:"mime_type"`
	FileSizeBytes int    `json:"file_size"`
}

type IncomingMessage struct {
	Text  *string
	Voice *Voice `json:"voice"`
	From  From   `json:"from"`
	Chat  Chat   `json:"chat"`
}
