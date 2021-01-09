package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	)

type Storage interface {
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{db: db}
}
