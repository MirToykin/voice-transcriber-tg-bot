-- name: GetUserByLogin :one
SELECT id, login, lang, created, updated FROM users WHERE login = ?;

-- name: CreateUser :exec
INSERT INTO users (login, lang)
values (?, ?);