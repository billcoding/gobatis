package gobatis

import (
	"database/sql"
	"time"
)

//default db config
var defaultDBConfig = &DBConfig{
	MaxIdleConns:    2,
	MaxOpenConns:    10,
	ConnMaxLifetime: 0,
}

//Define DB struct
type DB struct {
	db *sql.DB //built in db
}

//Define DBConfig struct
type DBConfig struct {
	MaxIdleConns    int           //Max idle conn num
	MaxOpenConns    int           //Max open conn num
	ConnMaxLifetime time.Duration //Conn max lifetime
}

//Begin a tx
func (db *DB) Begin() *Tx {
	tx, err := db.db.Begin()
	if err != nil {
		//batis.Error("tx begin err : %v", err)
		return nil
	}
	return &Tx{
		db: db.db,
		tx: tx,
	}
}
