package repository

import (
	"database/sql"
	"fmt"
)

func OpenAndPingDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		fmt.Printf("Something wrong with db %s\n", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Something wrong with db %s\n", err)
	}
	return db
}
