package gobatis

import (
	"database/sql"
	"strings"
	"text/template"
)

type UpdateMapper struct {
	gfuncMap     *template.FuncMap
	printSql     bool   //print sql
	logger       *log   //logger
	binding      string //binding key
	id           string //id
	tx           *TX    //sql tx
	db           *DB    //sql db
	originalSql  string //original sql
	sql          string //sql
	affectedRows int64  //affected rows
	insertedId   int64  //inserted id
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
	gfuncMap := joinFuncMap(*m.gfuncMap, funcMap)
	if len(gfuncMap) <= 0 {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(gfuncMap).Parse(m.originalSql))
	}
	var builder strings.Builder
	err := t.Execute(&builder, data)
	if err != nil {
		m.logger.Error(err.Error())
		return m
	}
	m.sql = builder.String()
	return m
}

// Exec update exec
func (m *UpdateMapper) Exec(args ...interface{}) error {
	return m.ExecWithParamsArgs(nil, args...)
}

// ExecWithParams update exec with named params
func (m *UpdateMapper) ExecWithParams(params ...*Param) error {
	return m.ExecWithParamsArgs(params)
}

//ExecWithParamsArgs update exec with named params
func (m *UpdateMapper) ExecWithParamsArgs(params []*Param, args ...interface{}) error {
	var result sql.Result
	var err error

	if params != nil {
		//replace namedParam
		m.replaceParams(params...)
	}

	if m.tx != nil {
		result, err = m.updateByTx(args...)
	} else {
		result, err = m.updateByDB(args...)
	}

	if m.printSql {
		m.logger.Info("[SQL]binding[%s] update[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, args)
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

func (m *UpdateMapper) replaceParams(params ...*Param) {
	m.sql = replaceParams(m.originalSql, params...)
}

func (m *UpdateMapper) updateByTx(args ...interface{}) (sql.Result, error) {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	if args != nil && len(args) > 0 {
		result, err := m.tx.tx.Exec(m.sql, args...)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
			return nil, err
		}
		return result, nil
	} else {
		result, err := m.tx.tx.Exec(m.sql)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] updateByTx error : %v", m.binding, m.id, err)
			return nil, err
		}
		return result, nil
	}
}

func (m *UpdateMapper) updateByDB(args ...interface{}) (sql.Result, error) {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	if args != nil && len(args) > 0 {
		result, err := m.db.db.Exec(m.sql, args...)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return result, nil
	} else {
		result, err := m.db.db.Exec(m.sql)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] updateByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return result, nil
	}
}
