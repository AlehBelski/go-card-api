package repository

import (
    "database/sql"
)

type DB struct {
    *sql.DB
}
