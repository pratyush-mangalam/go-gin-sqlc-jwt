package db

import (
	"database/sql"
	db "go-boiler-plate/db/sqlc"
)

// Store defines all functions to execute b queries and transactions.
type Store interface {
	db.Querier
}

// SOLStore provides all functions to execute SQL queries and transactions.

type SQLStore struct {
	db *sql.DB
	db.Querier
}

func NewStore(sqldb *sql.DB) Store {
	return &SQLStore{
		db:      sqldb,
		Querier: db.New(sqldb),
	}
}
