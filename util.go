package gobatis

import (
	"text/template"
)

func joinFuncMap(m1, m2 template.FuncMap) template.FuncMap {
	funcMap := make(template.FuncMap, 0)
	if m1 != nil && len(m1) > 0 {
		for k, v := range m1 {
			funcMap[k] = v
		}
	}
	if m2 != nil && len(m2) > 0 {
		for k, v := range m2 {
			funcMap[k] = v
		}
	}
	return funcMap
}
