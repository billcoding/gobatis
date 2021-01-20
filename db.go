package gobatis

import (
	"database/sql"
)

// DB struct
type DB struct {
	db *sql.DB //built in db
}

// Begin a tx
func (db *DB) Begin() *TX {
	tx, err := db.db.Begin()
	if err != nil {
		return nil
	}
	return &TX{
		db: db.db,
		tx: tx,
	}
}
