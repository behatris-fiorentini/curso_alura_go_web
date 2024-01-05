package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	connection := "user=root dbname=curso_alura_go_web password=secret host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}
