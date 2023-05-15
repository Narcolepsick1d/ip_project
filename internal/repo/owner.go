package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type OwnerPostgres struct {
	db *pgxpool.Pool
}

func NewOwnerPostgres(db *pgxpool.Pool) *OwnerPostgres {
	return &OwnerPostgres{db}
}
func (r *OwnerPostgres) Create(client_user_id int64) error {
	log.Println(client_user_id)
	_, err := r.db.Exec(context.Background(), "INSERT INTO owners (user_id) VALUES ($1)", client_user_id)
	log.Println(err)
	return err
}
