package gobatis

import (
	"fmt"
	"math"
)

// Page select page
func (m *SelectMapper) Page(rptr interface{}, offset, size int) *Page {
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
	if m.printSql {
		m.logger.Debugf("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, m.args)
	}
	//First query total count
	totalRows := m.queryCountByDB()
	if totalRows > 0 && offset >= 0 && offset <= totalRows-1 {
		page.TotalRows = totalRows
		page.TotalPage = int(math.Ceil(float64(totalRows) / float64(size)))
		// valid range
		m.extraSql = fmt.Sprintf(" limit %d,%d", offset, size)
		page.List = m.Exec().List(rptr)
		m.extraSql = ""
	}
	return page
}

// PageMap select Page
func (m *SelectMapper) PageMap(offset, size int) *PageMap {
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
	if m.printSql {
		m.logger.Debugf("binding[%s] selectPage[%s] exec : sql(%v), args(%v)", m.binding, m.id, m.sql, m.args)
	}
	//First query total count
	totalRows := m.queryCountByDB()
	if totalRows > 0 && offset >= 0 && offset <= totalRows-1 {
		page.TotalRows = totalRows
		page.TotalPage = int(math.Ceil(float64(totalRows) / float64(size)))
		// valid range
		m.extraSql = fmt.Sprintf(" limit %d,%d", offset, size)
		page.List = m.Exec().MapList()
		m.extraSql = ""
	}
	return page
}
