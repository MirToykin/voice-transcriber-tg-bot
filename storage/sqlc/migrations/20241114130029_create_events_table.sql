-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type INTEGER NOT NULL,
    file_path TEXT,
    file_size INTEGER,
    text TEXT,
    meta TEXT,
    processed BOOLEAN NOT NULL DEFAULT 0
);

CREATE INDEX idx_events_processed ON events (processed);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
