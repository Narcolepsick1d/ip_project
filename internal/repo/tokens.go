package repo

import (
	"awesomeProject2/internal/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Tokens struct {
	db *pgxpool.Pool
}

func NewTokens(db *pgxpool.Pool) *Tokens {
	return &Tokens{db}
}

func (r *Tokens) Create(token models.RefreshSession) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO refresh_tokens (user_id, token, expires_at) values ($1, $2, $3)",
		token.UserID, token.Token, token.ExpiresAt)

	return err
}

func (r *Tokens) Get(token string) (models.RefreshSession, error) {
	var t models.RefreshSession
	err := r.db.QueryRow(context.Background(), "SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token=$1", token).
		Scan(&t.ID, &t.UserID, &t.Token, &t.ExpiresAt)
	if err != nil {
		return t, err
	}

	_, err = r.db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE user_id=$1", t.UserID)

	return t, err
}
