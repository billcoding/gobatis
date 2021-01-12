package gobatis

import "database/sql"

// Tx struct
type Tx struct {
	db *sql.DB //built in db
	tx *sql.Tx //built in tx
}

func (tx *Tx) commit() error {
	return tx.tx.Commit()
}

func (tx *Tx) rollback() error {
	return tx.tx.Rollback()
}
