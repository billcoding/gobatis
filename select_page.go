package gobatis

import (
	"fmt"
	"math"
)

// Page select page
func (m *selectMapper) Page(rptr interface{}, offset, size int, args ...interface{}) *page {
	return m.PageWithParamsArgs(rptr, offset, size, nil, args...)
}

// PageWithParams select exec with params
func (m *selectMapper) PageWithParams(rptr interface{}, offset, size int, params ...*Param) *page {
	return m.PageWithParamsArgs(rptr, offset, size, params)
}

// PageWithParamsArgs select exec with args and named params
func (m *selectMapper) PageWithParamsArgs(rptr interface{}, offset, size int, params []*Param, args ...interface{}) *page {
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

	page := &page{
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

	if totalRows > 0 {
		page.TotalRows = totalRows
		page.TotalPage = int(math.Ceil(float64(totalRows) / float64(size)))
	}

	//valid range
	if offset >= 0 && offset <= page.TotalRows-1 {
		m.extraSql = fmt.Sprintf(" limit %d,%d", offset, size)
		page.List = m.Exec(args...).List(rptr)
		m.extraSql = ""
	}
	return page
}
