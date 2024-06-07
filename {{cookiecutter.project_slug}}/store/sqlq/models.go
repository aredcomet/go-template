// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlq

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthUser struct {
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

type AuthUserToken struct {
	ID        int64              `json:"id"`
	UserID    int64              `json:"user_id"`
	Token     string             `json:"token"`
	TokenType string             `json:"token_type"`
	IsActive  bool               `json:"is_active"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}