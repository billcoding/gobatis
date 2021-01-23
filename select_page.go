package gobatis

import (
	"fmt"
	"math"
)

// Page select Page
func (m *SelectMapper) Page(rptr interface{}, offset, size int, args ...interface{}) *Page {
	return m.PageWithParamsArgs(rptr, offset, size, nil, args...)
}

// PageWithParams select exec with params
func (m *SelectMapper) PageWithParams(rptr interface{}, offset, size int, params ...*Param) *Page {
	return m.PageWithParamsArgs(rptr, offset, size, params)
}

// PageWithParamsArgs select exec with args and named params
func (m *SelectMapper) PageWithParamsArgs(rptr interface{}, offset, size int, params []*Param, args ...interface{}) *Page {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()

	if size <= 0 {
		size = 10
	}

	if offset < 0 {
		offset = 0
	}

	page := &Page{
		Offset: offset,
		Size:   size,
		List:   []interface{}{},
	}

	if params != nil {
		//replace named param
		m.replaceParams(params...)
	}

	//First query total count
	totalRows := m.queryCountByDB(args...)

	if m.printSql {
		m.logger.Info("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, args)
	}

	if totalRows > 0 && offset >= 0 && offset <= totalRows-1 {
		page.TotalRows = totalRows
		page.TotalPage = int(math.Ceil(float64(totalRows) / float64(size)))
		// valid range
		m.extraSql = fmt.Sprintf(" limit %d,%d", offset, size)
		page.List = m.Exec(args...).List(rptr)
		m.extraSql = ""
	}
	return page
}

// PageMap select Page
func (m *SelectMapper) PageMap(offset, size int, args ...interface{}) *PageMap {
	return m.PageMapWithParamsArgs(offset, size, nil, args...)
}

// PageMapWithParams select exec with params
func (m *SelectMapper) PageMapWithParams(offset, size int, params ...*Param) *PageMap {
	return m.PageMapWithParamsArgs(offset, size, params)
}

// PageMapWithParamsArgs select exec with args and named params
func (m *SelectMapper) PageMapWithParamsArgs(offset, size int, params []*Param, args ...interface{}) *PageMap {
	defer func() {
		if re := recover(); re != nil {
			m.logger.Error("%v", re)
		}
	}()

	if size <= 0 {
		size = 10
	}

	if offset < 0 {
		offset = 0
	}

	page := &PageMap{
		Offset: offset,
		Size:   size,
		List:   []map[string]interface{}{},
	}

	if params != nil {
		//replace named param
		m.replaceParams(params...)
	}

	//First query total count
	totalRows := m.queryCountByDB(args...)

	if m.printSql {
		m.logger.Info("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, args)
	}

	if totalRows > 0 && offset >= 0 && offset <= totalRows-1 {
		page.TotalRows = totalRows
		page.TotalPage = int(math.Ceil(float64(totalRows) / float64(size)))
		// valid range
		m.extraSql = fmt.Sprintf(" limit %d,%d", offset, size)
		page.List = m.Exec(args...).MapList()
		m.extraSql = ""
	}
	return page
}
