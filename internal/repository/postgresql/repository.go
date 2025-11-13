package postgresql

import "database/sql"

type DBStorage struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DBStorage {
	return &DBStorage{
		db: db,
	}
}
