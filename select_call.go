package gobatis

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

type selectCall struct {
	sm     *SelectMapper
	args   []interface{}
	logger *logrus.Logger
	err    error
	rows   *sql.Rows
	rptr   interface{}
}

// Scan rows to dest
func (c *selectCall) Scan(dest ...interface{}) error {
	//fixed
	//close rows
	//release conn
	defer func() {
		err := c.rows.Close()
		if err != nil {
			c.logger.Panicf("%v", err)
		}
	}()
	if c.rows.Next() {
		err := c.rows.Scan(dest...)
		if err != nil {
			c.logger.Panicf("%v", err)
		}
	}
	return c.rows.Err()
}

// Single return single record
func (c *selectCall) Single() interface{} {
	var r interface{}
	err := c.Scan(&r)
	if err != nil {
		c.logger.Panicf("%v", err)
	}
	return r
}

// SingleInt return single record
func (c *selectCall) SingleInt() int64 {
	var r int64
	err := c.Scan(&r)
	if err != nil {
		c.logger.Panicf("%v", err)
	}
	return r
}

// SingleFloat return single record
func (c *selectCall) SingleFloat() float64 {
	var r float64
	err := c.Scan(&r)
	if err != nil {
		c.logger.Panicf("%v", err)
	}
	return r
}

// SingleString return single record
func (c *selectCall) SingleString() string {
	var r string
	err := c.Scan(&r)
	if err != nil {
		c.logger.Panicf("%v", err)
	}
	return r
}

// List get list rows
func (c *selectCall) List(ptr interface{}) []interface{} {
	c.rptr = ptr
	return c.scanStruct()
}

// MapList get map rows
func (c *selectCall) MapList() []map[string]interface{} {
	return c.scanMap()
}

// Call rows
func (c *selectCall) Call(callback func(rows *sql.Rows)) {
	if callback == nil {
		return
	}
	func() {
		defer func() {
			err := c.rows.Close()
			if err != nil {
				c.logger.Panicf("%v", err)
			}
		}()
		callback(c.rows)
	}()
}

func (c *selectCall) scanStruct() []interface{} {
	//must be kind of Ptr
	if reflect.TypeOf(c.rptr).Kind() != reflect.Ptr {
		c.logger.Panicf("structPtr must be the kind of reflect.Ptr")
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
			c.logger.Panicf("%v", err)
		}
	}()
	for c.rows.Next() {
		//make struct field num length slice
		fieldAdds := make([]interface{}, len(columns))
		//match's column and field
		//create new struct == dynamic create struct object
		nrv := reflect.New(rt).Elem()
		for i, column := range columns {
			if fieldName, have := fieldMap[column]; have {
				field := nrv.FieldByName(fieldName)
				//put field's address into fieldAdds
				fieldAdds[i] = field.Addr().Interface()
			} else {
				var unused interface{}
				fieldAdds[i] = &unused
			}
		}
		err := c.rows.Scan(fieldAdds...)
		if err != nil {
			c.logger.Panicf("%v", err)
		}
		list = append(list, nrv.Addr().Interface())
	}
	return list
}

func (c *selectCall) scanMap() []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	columns, _ := c.rows.Columns()
	//release conn
	defer func() {
		err := c.rows.Close()
		if err != nil {
			c.logger.Panicf("%v", err)
		}
	}()
	for c.rows.Next() {
		m := make(map[string]interface{}, 0)
		addrs := make([]interface{}, len(columns))
		for i := range columns {
			var obj string
			addrs[i] = &obj
		}
		err := c.rows.Scan(addrs...)
		if err != nil {
			c.logger.Panicf("%v", err)
		}
		for i, column := range columns {
			m[column] = getInterfaceVal(*(addrs[i].(*string)))
		}
		list = append(list, m)
	}
	return list
}

// TODO supports more types
func getInterfaceVal(val string) interface{} {
	// int64
	ival, err := strconv.ParseInt(val, 10, 64)
	if err == nil {
		return ival
	}

	// float64
	fval, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return fval
	}

	return val
}

func (c *selectCall) getFieldMap() map[string]string {
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
