package db

import (
	"AuthGoTest/pkg/e"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(connString string) (db *DB, err error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, e.Wrap("can't connect to db", err)
	}

	return &DB{pool: pool}, nil
}
