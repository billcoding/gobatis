package gobatis

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"strings"
	"text/template"
)

type UpdateMapper struct {
	funcMap      *template.FuncMap
	printSql     bool
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
	gfuncMap := joinFuncMap(*m.funcMap, funcMap)
	if len(gfuncMap) <= 0 {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(gfuncMap).Parse(m.originalSql))
	}
	var builder strings.Builder
	err := t.Execute(&builder, data)
	if err != nil {
		m.logger.Errorf(err.Error())
	}
	m.sql = builder.String()
	return m
}

// Params set params
func (m *UpdateMapper) Params(params ...*Param) *UpdateMapper {
	m.sql = replaceParams(m.originalSql, params...)
	return m
}

// Args set args
func (m *UpdateMapper) Args(args ...interface{}) *UpdateMapper {
	m.args = args
	return m
}

// Exec update exec
func (m *UpdateMapper) Exec() error {
	var result sql.Result
	var err error
	result, err = m.updateByDB()
	if m.printSql {
		m.logger.Infof("[SQL]binding[%s] update[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, m.args)
	}
	if err != nil {
		return err
	}
	if result != nil {
		ra, _ := result.RowsAffected()
		li, _ := result.LastInsertId()
		m.affectedRows = ra
		m.insertedId = li
	}
	return nil
}

func (m *UpdateMapper) updateByTx(tx *sql.Tx) (sql.Result, error) {
	if m.args != nil && len(m.args) > 0 {
		result, err := tx.Exec(m.sql, m.args...)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
		}
		return result, nil
	} else {
		result, err := tx.Exec(m.sql)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
		}
		return result, nil
	}
}

func (m *UpdateMapper) updateByDB() (sql.Result, error) {
	if m.args != nil && len(m.args) > 0 {
		result, err := m.db.db.Exec(m.sql, m.args...)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
		}
		return result, nil
	} else {
		result, err := m.db.db.Exec(m.sql)
		if err != nil {
			m.logger.Errorf("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
		}
		return result, nil
	}
}
