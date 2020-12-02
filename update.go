package gobatis

import (
	"database/sql"
	"strings"
	"text/template"
)

//Define update mapper struct
type updateMapper struct {
	binding      string //binding key
	id           string //id
	tx           *Tx    //sql tx
	db           *DB    //sql db
	originalSql  string //original sql
	sql          string //sql
	affectedRows int64  //affected rows
	insertedId   int64  //inserted id
}

//Get affectedRows
func (m *updateMapper) AffectedRows() int64 {
	return m.affectedRows
}

//Get insertedId
func (m *updateMapper) InsertedId() int64 {
	return m.insertedId
}

//Prepare using text/template
func (m *updateMapper) Prepare(data interface{}) *updateMapper {
	return m.PrepareWithFunc(data, nil)
}

//Prepare using text/template
func (m *updateMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *updateMapper {
	var t *template.Template
	if funcMap == nil {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(funcMap).Parse(m.originalSql))
	}
	var builder strings.Builder
	err := t.Execute(&builder, data)
	if err != nil {
		batis.Error(err.Error())
		return m
	}
	m.sql = builder.String()
	return m
}

//Update exec
func (m *updateMapper) Exec(args ...interface{}) error {
	return m.ExecWithParamsArgs(nil, args...)
}

//Update exec with named params
func (m *updateMapper) ExecWithParams(params ...*NamedParam) error {
	return m.ExecWithParamsArgs(params)
}

//Update exec with named params
func (m *updateMapper) ExecWithParamsArgs(params []*NamedParam, args ...interface{}) error {
	var result sql.Result
	var err error

	//replace named params
	m.sql = replaceNamedParams(m.sql, params...)

	if m.tx != nil {
		result, err = updateByTx(m, m.tx, m.sql, args...)
	} else {
		result, err = updateByDB(m, m.db, m.sql, args...)
	}

	if batis.showSql {
		batis.Info("binding[%s] update[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, args)
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
