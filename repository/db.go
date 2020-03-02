package repository

import (
    "database/sql"
    "fmt"
)

type DB struct {
    *sql.DB
}

// NewDB creates and returns a new database connection using passed username, password, host and db name.
func NewDB(userName, userPassword, host, dbName string) (*DB, error) {
    dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", userName, userPassword, host, dbName)
    db, err := sql.Open("postgres", dataSource)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &DB{db}, nil
}
