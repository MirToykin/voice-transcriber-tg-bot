package sqlite

type Event struct {
	ID        uint   `db:"id"`
	Username  string `db:"username"`
	FilePath  string `db:"file_path"`
	Processed bool   `db:"processed"`
}
