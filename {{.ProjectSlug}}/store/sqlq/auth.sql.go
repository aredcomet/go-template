// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth.sql

package sqlq

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getActiveRefreshToken = `-- name: GetActiveRefreshToken :one
SELECT aut.id, aut.user_id, aut.token, aut.token_type, aut.is_active, aut.expires_at, aut.created_at
FROM auth_user_token aut
WHERE aut.token = $1
  AND aut.token_type = 'refresh'
  AND aut.is_active = TRUE
LIMIT 1
`

func (q *Queries) GetActiveRefreshToken(ctx context.Context, token string) (AuthUserToken, error) {
	row := q.db.QueryRow(ctx, getActiveRefreshToken, token)
	var i AuthUserToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.TokenType,
		&i.IsActive,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT au.id,
       first_name,
       last_name,
       username,
       email,
       phone,
       '' as password,
       is_superuser,
       is_active,
       is_staff,
       is_bot,
       last_login,
       created_at,
       updated_at
FROM "auth_user" au
WHERE au.id = $1
LIMIT 1
`

type GetUserByIdRow struct {
	ID          int64              `json:"id"`
	FirstName   pgtype.Text        `json:"first_name"`
	LastName    pgtype.Text        `json:"last_name"`
	Username    string             `json:"username"`
	Email       string             `json:"email"`
	Phone       pgtype.Text        `json:"phone"`
	Password    string             `json:"password"`
	IsSuperuser bool               `json:"is_superuser"`
	IsActive    bool               `json:"is_active"`
	IsStaff     bool               `json:"is_staff"`
	IsBot       bool               `json:"is_bot"`
	LastLogin   pgtype.Timestamptz `json:"last_login"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetUserById(ctx context.Context, id int64) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.IsSuperuser,
		&i.IsActive,
		&i.IsStaff,
		&i.IsBot,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT au.id, au.first_name, au.last_name, au.username, au.email, au.phone, au.password, au.is_superuser, au.is_active, au.is_staff, au.is_bot, au.last_login, au.created_at, au.updated_at
FROM "auth_user" au
WHERE au.username = $1
LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (AuthUser, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.IsSuperuser,
		&i.IsActive,
		&i.IsStaff,
		&i.IsBot,
		&i.LastLogin,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const saveAuthTokenPair = `-- name: SaveAuthTokenPair :exec
INSERT INTO auth_user_token (user_id, token, token_type, expires_at, created_at)
VALUES ($1, $2, 'access', $3, $4),
       ($1, $5, 'refresh', $3, $4)
`

type SaveAuthTokenPairParams struct {
	UserID    int64              `json:"user_id"`
	Token     string             `json:"token"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Token_2   string             `json:"token_2"`
}

func (q *Queries) SaveAuthTokenPair(ctx context.Context, arg SaveAuthTokenPairParams) error {
	_, err := q.db.Exec(ctx, saveAuthTokenPair,
		arg.UserID,
		arg.Token,
		arg.ExpiresAt,
		arg.CreatedAt,
		arg.Token_2,
	)
	return err
}
