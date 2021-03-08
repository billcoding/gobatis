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
	args        []interface{}
}

// Prepare using text/template
func (m *SelectMapper) Prepare(data interface{}) *SelectMapper {
	return m.PrepareWithFunc(data, nil)
}

// PrepareWithFunc using text/template with func
func (m *SelectMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *SelectMapper {
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

// Args set args
func (m *SelectMapper) Args(args ...interface{}) *SelectMapper {
	m.args = args
	return m
}

// Params set params
func (m *SelectMapper) Params(params ...*Param) *SelectMapper {
	m.sql = replaceParams(m.originalSql, params...)
	return m
}

// Exec select exec
func (m *SelectMapper) Exec() *selectCall {
	var rows *sql.Rows
	var err error
	rows, err = m.queryByDB()
	if m.printSql {
		m.logger.Info("binding[%s] select[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql+m.extraSql, m.args)
	}
	if err != nil {
		return &selectCall{err: err}
	}
	return &selectCall{
		sm:     m,
		args:   m.args,
		logger: m.logger,
		rows:   rows,
	}
}

func (m *SelectMapper) queryCountByDB() int {
	rx := regexp.MustCompile(`^\s*[Ss][Ee][Ll][Ee][Cc][Tt]([\s\S]+)[Ff][Rr][Oo][Mm]`)
	csql := rx.ReplaceAllString(m.sql, " select count(*) from ")
	var rows *sql.Rows
	var err error
	if m.args != nil && len(m.args) > 0 {
		rows, err = m.db.db.Query(csql, m.args...)
	} else {
		rows, err = m.db.db.Query(csql)
	}
	if err != nil {
		m.logger.Error("binding[%s] select[%s] queryCountByDB error : %v", m.binding, m.id, err)
		panic(err)
	}
	if m.printSql {
		m.logger.Info("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, csql, m.args)
	}
	defer func() {
		_ = rows.Close()
	}()
	c := 0
	if rows.Next() {
		_ = rows.Scan(&c)
	}
	return c
}

func (m *SelectMapper) queryByDB() (*sql.Rows, error) {
	if m.args != nil && len(m.args) > 0 {
		rows, err := m.db.db.Query(m.sql+m.extraSql, m.args...)
		if err != nil {
			m.logger.Error("binding[%s] select[%s] queryByDB error : %v", m.binding, m.id, err)
			panic(err)
		}
		return rows, err
	} else {
		rows, err := m.db.db.Query(m.sql + m.extraSql)
		if err != nil {
			m.logger.Error("binding[%s] select[%s] queryByDB error : %v", m.binding, m.id, err)
			panic(err)
		}
		return rows, err
	}
}
