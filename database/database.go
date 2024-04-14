package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connStr := "user=*** dbname=loja password==*** host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
