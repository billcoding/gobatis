package gobatis

import (
	"database/sql"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
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
func replaceNamedParams(sql string, params ...*NamedParam) string {
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

//Return struct field name-column mapping
func getFieldMap(logger *log, structPtr interface{}) map[string]string {
	defer func() {
		if re := recover(); re != nil {
			logger.Error("%v", re)
		}
	}()
	reflectType := reflect.TypeOf(structPtr).Elem()
	fieldNum := reflectType.NumField()
	fieldMap := make(map[string]string, 0)
	for i := 0; i < fieldNum; i++ {
		field := reflectType.Field(i)
		if colName, ok := field.Tag.Lookup("db"); ok {
			fieldMap[colName] = field.Name
		} else {
			fieldMap[field.Name] = field.Name
		}
	}
	return fieldMap
}

//Scan a struct from rows
func scanStruct(logger *log, rows *sql.Rows, structPtr interface{}) []interface{} {
	defer func() {
		if re := recover(); re != nil {
			logger.Error("%v", re)
		}
	}()
	//must be kind of Ptr
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		logger.Error("structPtr must be the kind of reflect.Ptr")
	}

	//receive the struct type
	rt := reflect.TypeOf(structPtr).Elem()
	//get the struct field name map
	fieldMap := getFieldMap(logger, structPtr)
	//get rows's columns
	columns, _ := rows.Columns()
	//make return slice
	list := make([]interface{}, 0)

	//release conn
	defer rows.Close()
	for rows.Next() {
		//make struct field num length slice
		fieldAddrs := make([]interface{}, len(columns))
		//match's column and field
		//create new struct == dynamic create struct object
		nrv := reflect.New(rt).Elem()
		for i, column := range columns {
			if fieldName, have := fieldMap[column]; have {
				field := nrv.FieldByName(fieldName)
				//put field's address into fieldAddrs
				fieldAddrs[i] = field.Addr().Interface()
			} else {
				var unused interface{}
				fieldAddrs[i] = &unused
			}
		}
		rows.Scan(fieldAddrs...)
		list = append(list, nrv.Addr().Interface())
	}
	return list
}
