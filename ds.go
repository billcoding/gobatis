package gobatis

import (
	"database/sql"
	"time"
)

type multiDS struct {
	mds             map[string]*DS
	maxOpenConn     int
	maxIdleConn     int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
}

// DS struct
type DS struct {
	// Named of DS
	Name string
	// DSN of DS
	DSN string
	db  *DB
}

// DS select
func (m *mapper) DS(ds string) *mapper {
	mds, have := m.multiDS.mds[ds]
	if !have {
		m.logger.Panicf("mapper: choose DS[%s] not registered", ds)
	}
	m.currentDS = mds
	m.logger.Debugf("mapper: choose DS[%s]", ds)
	return m
}

// Size of MultiDS
func (m *multiDS) Size() int {
	return len(m.mds)
}

// Add DS
func (m *multiDS) Add(name, dsn string) *DS {
	return m.AddWithDialect(name, dsn, MySQL)
}

// AddWithDialect add DS with dialect
func (m *multiDS) AddWithDialect(name, dsn string, dialect Dialect) *DS {
	db, err := sql.Open(string(dialect), dsn)
	if err != nil {
		panic(err)
	}
	{
		db.SetMaxOpenConns(m.maxOpenConn)
		db.SetMaxIdleConns(m.maxIdleConn)
		db.SetConnMaxLifetime(m.connMaxLifetime)
		db.SetConnMaxIdleTime(m.connMaxIdleTime)
	}
	ds := &DS{
		Name: name,
		DSN:  dsn,
		db: &DB{
			db: db,
		},
	}
	m.mds[name] = ds
	return ds
}

func (m *multiDS) defaultDS() (string, *DS) {
	if len(m.mds) <= 0 {
		panic("MultiDS: MultiDS is empty")
	}
	mds, have := m.mds["master"]
	if have {
		return "master", mds
	}
	ds := ""
	for name := range m.mds {
		ds = name
		break
	}
	return ds, m.mds[ds]
}
