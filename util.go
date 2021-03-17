package gobatis

import (
	"strings"
	"text/template"
)

func replaceParams(sql string, params ...*Param) string {
	replacedSql := sql
	if params != nil && len(params) > 0 {
		for _, param := range params {
			replace := param.replace
			val := param.val
			if strings.Contains(replacedSql, replace) {
				replacedSql = strings.ReplaceAll(replacedSql, replace, val)
			}
		}
	}
	return replacedSql
}

func joinFuncMap(m1, m2 template.FuncMap) template.FuncMap {
	gfuncMap := make(template.FuncMap, 0)
	if m1 != nil && len(m1) > 0 {
		for k, v := range m1 {
			gfuncMap[k] = v
		}
	}
	if m2 != nil && len(m2) > 0 {
		for k, v := range m2 {
			gfuncMap[k] = v
		}
	}
	return gfuncMap
}
