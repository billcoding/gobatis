package gobatis

import (
	"database/sql"
)

// DB struct
type DB struct {
	db *sql.DB //built in db
}

// Begin a tx
func (db *DB) Begin() *Tx {
	tx, err := db.db.Begin()
	if err != nil {
		return nil
	}
	return &Tx{
		db: db.db,
		tx: tx,
	}
}
