package data

import (
	"github.com/jmoiron/sqlx"
)

type DBStore struct {
	db *sqlx.DB
}

func Load() *sqlx.DB {
	db, err := sqlx.Open("postgres", "dbname=learn user=postgres password=x3127106 sslmode=disble")
	if err != nil {
		panic(err)
	}
	return db
}
