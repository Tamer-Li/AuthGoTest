package db

import (
	"AuthGoTest/internal/e"
	"database/sql"
)

type DB struct {
	pool *sql.DB
}

func Connect(urlDB string) (db *DB, err error) {
	defer func() { err = e.Wrap("can't connect to db", err) }()
	conn, err := sql.Open("postgres", urlDB)
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{pool: conn}, nil
}
