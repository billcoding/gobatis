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

func (tx *TX) Update(m *UpdateMapper) error {
	result, err := m.updateByTx(tx.tx)
	if err != nil {
		return err
	}
	m.insertedId, _ = result.LastInsertId()
	m.affectedRows, _ = result.RowsAffected()
	return nil
}

func (tx *TX) Commit() error {
	if err := tx.tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (tx *TX) Rollback() error {
	if err := tx.tx.Rollback(); err != nil {
		return err
	}
	return nil
}
