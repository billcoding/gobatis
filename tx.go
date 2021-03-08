package gobatis

import (
	"database/sql"
)

// TX struct
type TX struct {
	// built in tx
	tx *sql.Tx
}

func (b *Batis) Begin() *TX {
	_, d := b.MultiDS.defaultDS()
	if d == nil {
		panic("Required a default DS")
	}
	tx, err := d.db.db.Begin()
	if err != nil {
		panic(err)
	}
	return &TX{
		tx: tx,
	}
}

func (tx *TX) Update(m *UpdateMapper) {
	result, err := m.updateByTx(tx.tx)
	if err != nil {
		panic(err)
	}
	m.insertedId, _ = result.LastInsertId()
	m.affectedRows, _ = result.RowsAffected()
}

func (tx *TX) Commit() {
	if err := tx.tx.Commit(); err != nil {
		panic(err)
	}
}

func (tx *TX) Rollback() {
	if err := tx.tx.Rollback(); err != nil {
		panic(err)
	}
}
