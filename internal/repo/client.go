package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type ClientPostgres struct {
	db *pgxpool.Pool
}

func NewClientPostgres(db *pgxpool.Pool) *ClientPostgres {
	return &ClientPostgres{db}
}
func (r *ClientPostgres) Create(client_user_id int64) error {
	log.Println(client_user_id)
	_, err := r.db.Exec(context.Background(), "INSERT INTO clients (user_id) VALUES ($1)", client_user_id)
	log.Println(err)
	return err
}
