package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var counts int64

func NewPostgresConnection() (*pgxpool.Pool, error) {
	ctx := context.Background()
	for {
		db, err := pgxpool.Connect(ctx, "host=localhost port=5432 user=postgres dbname=ip_project sslmode=disable password=12345")
		if err != nil {
			log.Println("Postgres not yet ready ...")
			counts++

		} else {
			log.Println("Connected to Postgres!")
			return db, nil
		}
		if err := db.Ping(ctx); err != nil {
			return nil, err
		}
		if counts > 10 {
			log.Println(err)
			return nil, err
		}
		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue

	}
}
