package postgresql

import "database/sql"

type DBrepo struct {
	db *sql.DB
}

func NewDB(db *sql.DB) *DBrepo {
	return &DBrepo{
		db: db,
	}
}
