package gobatis

import (
	"database/sql"
	"regexp"
	"strings"
	"text/template"
)

type SelectMapper struct {
	gfuncMap    *template.FuncMap
	logger      *log
	printSql    bool
	binding     string
	id          string
	db          *DB
	originalSql string
	sql         string
	extraSql    string
}

// Prepare using text/template
func (m *SelectMapper) Prepare(data interface{}) *SelectMapper {
	return m.PrepareWithFunc(data, nil)
}

// PrepareWithFunc using text/template with func
func (m *SelectMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *SelectMapper {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
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

// Exec select exec
func (m *SelectMapper) Exec(args ...interface{}) *selectCall {
	return m.ExecWithParamsArgs(nil, args...)
}

// ExecWithParams select exec with params
func (m *SelectMapper) ExecWithParams(params ...*Param) *selectCall {
	return m.ExecWithParamsArgs(params)
}

// ExecWithParamsArgs select exec with args and named params
func (m *SelectMapper) ExecWithParamsArgs(params []*Param, args ...interface{}) *selectCall {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	var rows *sql.Rows
	var err error

	if params != nil {
		//replace namedParam
		m.replaceParams(params...)
	}

	rows, err = m.queryByDB(args...)

	if m.printSql {
		m.logger.Info("binding[%s] select[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql+m.extraSql, args)
	}

	if err != nil {
		return &selectCall{err: err}
	}

	return &selectCall{
		sm:     m,
		args:   args,
		logger: m.logger,
		rows:   rows,
	}
}

func (m *SelectMapper) replaceParams(params ...*Param) {
	m.sql = replaceParams(m.originalSql, params...)
}

func (m *SelectMapper) queryCountByDB(args ...interface{}) int {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()

	rx := regexp.MustCompile(`^\s*[Ss][Ee][Ll][Ee][Cc][Tt]([\s\S]+)[Ff][Rr][Oo][Mm]`)
	csql := rx.ReplaceAllString(m.sql, " select count(*) from ")

	var rows *sql.Rows
	var err error
	if args != nil && len(args) > 0 {
		rows, err = m.db.db.Query(csql, args...)
	} else {
		rows, err = m.db.db.Query(csql)
	}
	if err != nil {
		m.logger.Error("binding[%s] select[%s] queryCountByDB error : %v", m.binding, m.id, err)
		return 0
	}
	if m.printSql {
		m.logger.Info("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, csql, args)
	}
	defer rows.Close()
	c := 0
	if rows.Next() {
		rows.Scan(&c)
	}
	return c
}

func (m *SelectMapper) queryByDB(args ...interface{}) (*sql.Rows, error) {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	if args != nil && len(args) > 0 {
		rows, err := m.db.db.Query(m.sql+m.extraSql, args...)
		if err != nil {
			m.logger.Error("binding[%s] select[%s] queryByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return rows, err
	} else {
		rows, err := m.db.db.Query(m.sql + m.extraSql)
		if err != nil {
			m.logger.Error("binding[%s] select[%s] queryByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return rows, err
	}
}
