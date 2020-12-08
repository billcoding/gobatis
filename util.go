package gobatis

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//Get mapperNode files
func getMapperFiles(dir ...string) []string {
	var mapperFiles []string
	for _, mapperPath := range dir {
		fs, err := ioutil.ReadDir(mapperPath)
		if err != nil {
			continue
		}
		for _, f := range fs {
			curr := mapperPath + string(filepath.Separator) + f.Name()
			if f.IsDir() {
				mapperFiles = append(mapperFiles, getMapperFiles(curr)...)
				continue
			}
			match, err := regexp.MatchString("^.+\\.[xX][mM][lL]$", f.Name())
			if err == nil && match {
				mapperFiles = append(mapperFiles, curr)
			}
		}
	}
	return mapperFiles
}

//Prepare stmt
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
