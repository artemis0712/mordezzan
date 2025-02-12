// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package db

import (
	"context"
	"time"
)

const createSession = `-- name: CreateSession :one
INSERT INTO
    sessions (token, user_id, expires_at)
VALUES
    (?, ?, ?) RETURNING token, user_id, expires_at
`

type CreateSessionParams struct {
	Token     string    `json:"token"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.Token, arg.UserID, arg.ExpiresAt)
	var i Session
	err := row.Scan(&i.Token, &i.UserID, &i.ExpiresAt)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (username, email, password_hash)
VALUES
    (?, ?, ?) RETURNING id, username, email, password_hash, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE
    token = ?
`

func (q *Queries) DeleteSession(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteSession, token)
	return err
}

const getSession = `-- name: GetSession :one
SELECT
    s.token, s.user_id, s.expires_at,
    u.username,
    u.email
FROM
    sessions s
    JOIN users u ON s.user_id = u.id
WHERE
    s.token = ?
    AND s.expires_at > CURRENT_TIMESTAMP
LIMIT
    1
`

type GetSessionRow struct {
	Token     string    `json:"token"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
}

func (q *Queries) GetSession(ctx context.Context, token string) (GetSessionRow, error) {
	row := q.db.QueryRowContext(ctx, getSession, token)
	var i GetSessionRow
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.ExpiresAt,
		&i.Username,
		&i.Email,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id, username, email, password_hash, created_at
FROM
    users
WHERE
    email = ?
LIMIT
    1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT
    id, username, email, password_hash, created_at
FROM
    users
WHERE
    username = ?
LIMIT
    1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}
