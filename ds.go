package gobatis

import "database/sql"

//Define MultiDS type
type MultiDS map[string]*DS

//Define ds struct
type DS struct {
	Name   string //Named of DS
	DSN    string //DSN for DS
	db     *DB    //sql db
	Config *DBConfig
}

//Select ds
func (m *mapper) DS(ds string) *mapper {
	mds, have := m.multiDS[ds]
	if !have {
		m.logger.Error("[Mapper]Choose DS[%s] fail: not registered", ds)
		return m
	}
	m.currentDS = mds
	m.logger.Info("[Mapper]Choose DS[%s]", ds)
	return m
}

//Get DS size
func (m MultiDS) Size() int {
	return len(m)
}

//Add datasource
func (m MultiDS) Add(name, dsn string) *DS {
	return m.AddWithDialect(MySQL, name, dsn)
}

//Add datasource with dialect
func (m MultiDS) AddWithDialect(dialect Dialect, name, dsn string) *DS {
	db, err := sql.Open(string(dialect), dsn)
	if err != nil {
		panic(err)
	}
	ds := &DS{
		Name: name,
		DSN:  dsn,
		db: &DB{
			db: db,
		},
	}
	m[name] = ds
	return ds
}

//Get master ds
func (m MultiDS) defaultDS() (string, *DS) {
	if len(m) <= 0 {
		panic("[MultiDS]MultiDS is empty")
	}
	mds, have := m["master"]
	if have {
		return "master", mds
	}
	ds := ""
	for name := range m {
		ds = name
		break
	}
	return ds, m[ds]
}
