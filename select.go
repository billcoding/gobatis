package gobatis

import (
	"database/sql"
	"strings"
	"text/template"
)

//Define select mapper struct
type selectMapper struct {
	logger      *log   //logger
	printSql    bool   //print sql
	binding     string //binding key
	id          string //id
	db          *DB    //db
	originalSql string //original sql
	sql         string //sql
}

//Define selectCall struct
type selectCall struct {
	logger *log
	rows   *sql.Rows
}

//Get single rows
func (c *selectCall) Single(dists ...interface{}) error {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	//fixed
	//close rows
	//release conn
	defer c.rows.Close()
	if c.rows.Next() {
		c.rows.Scan(dists...)
	}
	return c.rows.Err()
}

//Get list rows
func (c *selectCall) List(structPtr interface{}) []interface{} {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	return scanStruct(c.logger, c.rows, structPtr)
}

//Call rows
func (c *selectCall) Call(callback func(rows *sql.Rows)) {
	if callback == nil {
		return
	}
	func() {
		defer func() {
			if re := recover(); re != nil {
				c.logger.Error("%v", re)
			}
		}()
		defer c.rows.Close()
		callback(c.rows)
	}()
}

//Prepare using text/template
func (m *selectMapper) Prepare(data interface{}) *selectMapper {
	return m.PrepareWithFunc(data, nil)
}

//Prepare using text/template
func (m *selectMapper) PrepareWithFunc(data interface{}, funcMap template.FuncMap) *selectMapper {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	var t *template.Template
	if funcMap == nil {
		t = template.Must(template.New("").Parse(m.originalSql))
	} else {
		t = template.Must(template.New("").Funcs(funcMap).Parse(m.originalSql))
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
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	var rows *sql.Rows
	var err error

	//replace namedParam
	m.sql = replaceNamedParams(m.sql, params...)
	rows, err = m.queryByDB(args...)

	if m.printSql {
		m.logger.Info("binding[%s] select[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, args)
	}

	if err != nil {
		return nil
	}

	return &selectCall{
		logger: m.logger,
		rows:   rows,
	}
}

//Query on db
func (m *selectMapper) queryByDB(args ...interface{}) (*sql.Rows, error) {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()
	if args != nil && len(args) > 0 {
		rows, err := m.db.db.Query(m.sql, args...)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] queryByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return rows, err
	} else {
		rows, err := m.db.db.Query(m.sql)
		if err != nil {
			m.logger.Error("binding[%s] update[%s] queryByDB error : %v", m.binding, m.id, err)
			return nil, err
		}
		return rows, err
	}
}
