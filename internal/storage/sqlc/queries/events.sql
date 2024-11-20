-- name: FetchUnprocessedEvents :many
SELECT id, type, file_path, file_size, text, meta FROM events WHERE processed = 0 LIMIT ?;

-- name: SaveUnprocessedEvent :exec
INSERT INTO events (type, file_path, file_size, text, meta, processed)
values (?, ?, ?, ?, ?, 0);

-- name: SetEventProcessed :exec
UPDATE events SET processed = 1 WHERE id = ?;

-- name: DeleteProcessedEvents :exec
DELETE FROM events WHERE processed = 1;