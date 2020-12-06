package gobatis

import (
	"database/sql"
	"reflect"
)

//Define selectCall struct
type selectCall struct {
	sm     *selectMapper //The select mapper
	args   []interface{} //The args
	logger *log          //The logger
	rows   *sql.Rows     //The rows
	rptr   interface{}   //The return type ptr
}

//Scan rows to dists
func (c *selectCall) Scan(dists ...interface{}) error {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	//fixed
	//close rows
	//release conn
	defer func() {
		err := c.rows.Close()
		if err != nil {
			c.logger.Error("%v", err)
		}
	}()
	if c.rows.Next() {
		err := c.rows.Scan(dists...)
		if err != nil {
			c.logger.Error("%v", err)
		}
	}
	return c.rows.Err()
}

//Return single record
func (c *selectCall) Single() interface{} {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	var r interface{}
	err := c.Scan(&r)
	if err != nil {
		c.logger.Error("%v", err)
	}
	return r
}

//Return single record
func (c *selectCall) SingleInt() int64 {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	var r int64
	err := c.Scan(&r)
	if err != nil {
		c.logger.Error("%v", err)
	}
	return r
}

//Return single record
func (c *selectCall) SingleFloat() float64 {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	var r float64
	err := c.Scan(&r)
	if err != nil {
		c.logger.Error("%v", err)
	}
	return r
}

//Return single record
func (c *selectCall) SingleString() string {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	var r string
	err := c.Scan(&r)
	if err != nil {
		c.logger.Error("%v", err)
	}
	return r
}

//Get list rows
func (c *selectCall) List(rptr interface{}) []interface{} {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	c.rptr = rptr
	return c.scanStruct()
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
		defer func() {
			err := c.rows.Close()
			if err != nil {
				c.logger.Error("%v", err)
			}
		}()
		callback(c.rows)
	}()
}

//Scan a struct from rows
func (c *selectCall) scanStruct() []interface{} {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	//must be kind of Ptr
	if reflect.TypeOf(c.rptr).Kind() != reflect.Ptr {
		c.logger.Error("structPtr must be the kind of reflect.Ptr")
	}

	//receive the struct type
	rt := reflect.TypeOf(c.rptr).Elem()
	//get the struct field name map
	fieldMap := c.getFieldMap()
	//get rows's columns
	columns, _ := c.rows.Columns()
	//make return slice
	list := make([]interface{}, 0)

	//release conn
	defer func() {
		err := c.rows.Close()
		if err != nil {
			c.logger.Error("%v", err)
		}
	}()
	for c.rows.Next() {
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
		err := c.rows.Scan(fieldAddrs...)
		if err != nil {
			c.logger.Error("%v", err)
		}
		list = append(list, nrv.Addr().Interface())
	}
	return list
}

//Return struct field name-column mapping
func (c *selectCall) getFieldMap() map[string]string {
	defer func() {
		if re := recover(); re != nil {
			c.logger.Error("%v", re)
		}
	}()
	reflectType := reflect.TypeOf(c.rptr).Elem()
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
