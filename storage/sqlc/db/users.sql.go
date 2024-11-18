// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (login, lang)
values (?, ?)
`

type CreateUserParams struct {
	Login string
	Lang  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Login, arg.Lang)
	return err
}

const getUserByLogin = `-- name: GetUserByLogin :one
SELECT id, login, lang, created, updated FROM users WHERE login = ?
`

func (q *Queries) GetUserByLogin(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByLogin, login)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Login,
		&i.Lang,
		&i.Created,
		&i.Updated,
	)
	return i, err
}
