package gobatis

import (
	"database/sql"
	"time"
)

//Define DB struct
type DB struct {
	db *sql.DB //built in db
}

//Define DBConfig struct
type DBConfig struct {
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

//Begin a tx
func (db *DB) Begin() *Tx {
	tx, err := db.db.Begin()
	if err != nil {
		batis.LogFatal("tx begin err : %v", err)
		return nil
	}
	return &Tx{
		db: db.db,
		tx: tx,
	}
}
