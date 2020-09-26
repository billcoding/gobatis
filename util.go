package gobatis

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

var namedParamNotSetErr = errors.New("named param not set fully")

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

//return struct field name-column mapping
func getFieldMap(structPtr interface{}) map[string]string {
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

//scan a struct from rows
func scanStruct(rows *sql.Rows, structPtr interface{}) []interface{} {
	//must be kind of Ptr
	if reflect.TypeOf(structPtr).Kind() != reflect.Ptr {
		panic("structPtr must be the kind of reflect.Ptr")
	}

	//receive the struct type
	rt := reflect.TypeOf(structPtr).Elem()
	//get the struct field name map
	fieldMap := getFieldMap(structPtr)
	//get rows's columns
	columns, _ := rows.Columns()
	//make return slice
	list := make([]interface{}, 0)

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

	//release conn
	rows.Close()
	return list
}

func queryByTx(tx *Tx, sql string, args ...interface{}) *sql.Rows {
	if args != nil && len(args) > 0 {
		rows, err := tx.tx.Query(sql, args...)
		if err != nil {
			batis.LogFatal("queryByTx error : %v", err)
			return nil
		}
		return rows
	} else {
		rows, err := tx.tx.Query(sql)
		if err != nil {
			batis.LogFatal("queryByTx error : %v", err)
			return nil
		}
		return rows
	}
}

func queryByDB(db *DB, sql string, args ...interface{}) *sql.Rows {
	if args != nil && len(args) > 0 {
		rows, err := db.db.Query(sql, args...)
		if err != nil {
			batis.LogFatal("queryByDB error : %v", err)
			return nil
		}
		return rows
	} else {
		rows, err := db.db.Query(sql)
		if err != nil {
			batis.LogFatal("queryByDB error : %v", err)
			return nil
		}
		return rows
	}
}

func updateByTx(tx *Tx, sql string, args ...interface{}) sql.Result {
	if args != nil && len(args) > 0 {
		result, err := tx.tx.Exec(sql, args...)
		if err != nil {
			batis.LogFatal("updateByTx error : %v", err)
			return nil
		}
		return result
	} else {
		result, err := tx.tx.Exec(sql)
		if err != nil {
			batis.LogFatal("updateByTx error : %v", err)
			return nil
		}
		return result
	}
}

func updateByDB(db *DB, sql string, args ...interface{}) sql.Result {
	if args != nil && len(args) > 0 {
		result, err := db.db.Exec(sql, args...)
		if err != nil {
			batis.LogFatal("updateByDB error : %v", err)
			return nil
		}
		return result
	} else {
		result, err := db.db.Exec(sql)
		if err != nil {
			batis.LogFatal("updateByDB error : %v", err)
			return nil
		}
		return result
	}
}
