package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/app?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
