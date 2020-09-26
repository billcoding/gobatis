package gobatis

import (
	"database/sql"
	"strings"
	"text/template"
)

//Define select mapper struct
type selectMapper struct {
	id          string //id
	tx          *Tx    //tx
	db          *DB    //db
	originalSql string //original sql
	sql         string //sql
}

//Define selectCall struct
type selectCall struct {
	rows *sql.Rows
}

//Get single rows
func (call *selectCall) Single(dists ...interface{}) error {
	if call.rows.Next() {
		call.rows.Scan(dists...)
	}
	//fixed
	//close rows
	//release conn
	call.rows.Close()
	return call.rows.Err()
}

//Get list rows
func (call *selectCall) List(structPtr interface{}) []interface{} {
	return scanStruct(call.rows, structPtr)
}

//Call rows
func (call *selectCall) Call(callback func(rows *sql.Rows)) {
	if callback != nil {
		callback(call.rows)
	}
}

//Prepare using text/template
func (m *selectMapper) Prepare(data interface{}) *selectMapper {
	return m.PrepareWithFunc(data, nil)
}

//Prepare using text/template
func (m *selectMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *selectMapper {
	var t *template.Template
	if funcMap == nil {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(funcMap).Parse(m.originalSql))
	}
	var builder strings.Builder
	err := t.Execute(&builder, data)
	if err != nil {
		batis.LogFatal(err.Error())
		return m
	}
	m.sql = builder.String()
	return m
}

//Select exec
func (m *selectMapper) Exec(args ...interface{}) *selectCall {
	return m.ExecWithParamsArgs(nil, args...)
}

//Select exec with params
func (m *selectMapper) ExecWithParams(params ...*NamedParam) *selectCall {
	return m.ExecWithParamsArgs(params)
}

//Select exec with args and named params
func (m *selectMapper) ExecWithParamsArgs(params []*NamedParam, args ...interface{}) *selectCall {
	var rows *sql.Rows
	//replace namedParam
	m.sql = replaceNamedParams(m.sql, params...)
	if m.tx != nil {
		rows = queryByTx(m.tx, m.sql, args...)
	} else {
		rows = queryByDB(m.db, m.sql, args...)
	}
	if batis.showSql {
		batis.LogInfo("exec [%v] : sql(%v) , args(%v)", m.id, m.sql, args)
	}
	return &selectCall{rows: rows}
}
