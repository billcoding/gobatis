package gobatis

import "database/sql"

type sqlQuery struct {
	db  *DB
	sql string
}

// SqlQuery return sql query from batis
func (b *Batis) SqlQuery(sql string) *sqlQuery {
	_, DS := b.MultiDS.defaultDS()
	return &sqlQuery{
		db:  DS.db,
		sql: sql,
	}
}

// SqlQuery return sql query from mapper
func (m *mapper) SqlQuery(sql string) *sqlQuery {
	return &sqlQuery{
		db:  m.currentDS.db,
		sql: sql,
	}
}

// Query query
func (s *sqlQuery) Query(args ...interface{}) (*sql.Rows, error) {
	return s.db.db.Query(s.sql, args...)
}

// Exec exec
func (s *sqlQuery) Exec(args ...interface{}) (sql.Result, error) {
	return s.db.db.Exec(s.sql, args...)
}
