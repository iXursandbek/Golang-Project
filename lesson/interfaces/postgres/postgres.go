package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Load() *sql.DB {
	db, err := sql.Open("postgres",
		`user=postgres 
		 dbname=joins 
		 sslmode=disable
	`)
	if err != nil {
		panic(err)
	}

	return db
}
