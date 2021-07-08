package gobatis

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"text/template"
)

type UpdateMapper struct {
	mu           *sync.Mutex
	funcMap      *template.FuncMap
	logger       *logrus.Logger
	binding      string
	id           string
	db           *DB
	originalSql  string
	sql          string
	affectedRows int64
	insertedId   int64
	args         []interface{}
}

// AffectedRows get affectedRows
func (m *UpdateMapper) AffectedRows() int64 {
	return m.affectedRows
}

// InsertedId get insertedId
func (m *UpdateMapper) InsertedId() int64 {
	return m.insertedId
}

// Prepare using text/template
func (m *UpdateMapper) Prepare(data interface{}) *UpdateMapper {
	return m.PrepareWithFunc(data, nil)
}

// PrepareWithFunc using text/template
func (m *UpdateMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *UpdateMapper {
	var t *template.Template
	funcMap2 := joinFuncMap(*m.funcMap, funcMap)
	if len(funcMap) <= 0 {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(funcMap2).Parse(m.originalSql))
	}
	var builder strings.Builder
	err := t.Execute(&builder, data)
	if err != nil {
		m.logger.Panicf(err.Error())
	}
	m.sql = builder.String()
	return m
}

// Args set args
func (m *UpdateMapper) Args(args ...interface{}) *UpdateMapper {
	m.args = args
	return m
}

// Exec update exec
func (m *UpdateMapper) Exec() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	var result sql.Result
	var err error
	result, err = m.updateByDB()
	m.logger.Debugf("SQL: binding[%s] update[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, m.args)
	if err != nil {
		return err
	}
	if result != nil {
		ra, _ := result.RowsAffected()
		li, _ := result.LastInsertId()
		m.affectedRows = ra
		m.insertedId = li
	}
	return err
}

func (m *UpdateMapper) updateByTx(tx *sql.Tx) (sql.Result, error) {
	if m.args != nil && len(m.args) > 0 {
		result, err := tx.Exec(m.sql, m.args...)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
		}
		return result, err
	} else {
		result, err := tx.Exec(m.sql)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
		}
		return result, err
	}
}

func (m *UpdateMapper) updateByDB() (sql.Result, error) {
	if m.args != nil && len(m.args) > 0 {
		result, err := m.db.db.Exec(m.sql, m.args...)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
		}
		return result, err
	} else {
		result, err := m.db.db.Exec(m.sql)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
		}
		return result, err
	}
}
