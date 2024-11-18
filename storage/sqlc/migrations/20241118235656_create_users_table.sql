-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        login VARCHAR(50) NOT NULL UNIQUE,
                        lang VARCHAR(10) NOT NULL,
                        created DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
                        updated DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TRIGGER set_updated_timestamp
    AFTER UPDATE ON users
    FOR EACH ROW
BEGIN
    UPDATE users
    SET updated = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_updated_timestamp;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
