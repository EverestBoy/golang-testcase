package database

import "database/sql"

type Database interface {
	GetDatabase() (*sql.DB, error)
}
