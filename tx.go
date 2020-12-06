package gobatis

import "database/sql"

//Define Tx struct
type Tx struct {
	db *sql.DB //built in db
	tx *sql.Tx //built in tx
}

//Commit the tx
func (tx *Tx) commit() error {

	return tx.tx.Commit()
}

//Rollback the tx
func (tx *Tx) rollback() error {
	return tx.tx.Rollback()
}
