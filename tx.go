package gobatis

import "database/sql"

// TX struct
type TX struct {
	db *sql.DB //built in db
	tx *sql.Tx //built in tx
}

func (tx *TX) commit() error {
	return tx.tx.Commit()
}

func (tx *TX) rollback() error {
	return tx.tx.Rollback()
}
