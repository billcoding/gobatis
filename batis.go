package gobatis

import (
	"database/sql"
	"sync"
)

var batis *Batis

//Define batis struct
type Batis struct {
	mutex             sync.Mutex            //mutex
	log               log                   //log
	showSql           bool                  //show sql
	dialect           Dialect               //choose db dialect
	dss               map[string]*ds        //multiple datasource
	dbConfig          *DBConfig             //dbConfig
	mapperPaths       []string              //mapper path
	parsedMapperPaths []string              //parsed mapper path
	mapperFiles       []string              //mapperNode files
	mappers           map[string]mapper     //mapper
	mapperNodes       map[string]mapperNode //mapper nodes
}

//new batis
func newBatis() *Batis {
	return &Batis{
		mutex:   sync.Mutex{},
		log:     log{},
		showSql: false,
		dialect: MySQL,
		dss:     map[string]*ds{},
		dbConfig: &DBConfig{
			MaxIdleConns:    2,
			MaxOpenConns:    10,
			ConnMaxLifetime: 0,
		},
		mapperPaths: []string{"./mapper"},
		mappers:     map[string]mapper{},
		mapperNodes: map[string]mapperNode{},
	}
}

//New batis
func init() {
	batis = newBatis()
}

//Return batis
func Default() *Batis {
	return batis
}

//Return new batis
func New() *Batis {
	return newBatis()
}

//Init batis
func (b *Batis) Init() *Batis {
	b.parseEnv()
	b.parseMapperPaths()
	b.scanMapper()
	return b
}

//Set showSql
func (b *Batis) ShowSql(showSql bool) *Batis {
	b.showSql = showSql
	return b
}

//Set dbConfig
func (b *Batis) DBConfig(dbConfig *DBConfig) *Batis {
	b.dbConfig = dbConfig
	return b
}

//Register datasource
func (b *Batis) RegisterDS(name, dsn string) *Batis {
	return b.RegisterDSWithConfig(name, dsn, nil)
}

//Register datasource with config
func (b *Batis) RegisterDSWithConfig(name, dsn string, dbConfig *DBConfig) *Batis {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	db, err := sql.Open(string(b.dialect), dsn)
	if err != nil {
		panic(err)
	}

	if dbConfig == nil {
		dbConfig = b.dbConfig
	}
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	b.dss[name] = &ds{
		dsn: dsn,
		db: &DB{
			db: db,
		},
	}
	return b
}

//Start scan mapper file for binding
func (b *Batis) scanMapper() *Batis {
	if len(b.parsedMapperPaths) <= 0 {
		return b
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	//collect mapperNode files
	for _, mapperPath := range b.parsedMapperPaths {
		b.LogInfo("collect mapper files : %v", mapperPath)
		for _, mf := range getMapperFiles(mapperPath) {
			b.mapperFiles = append(b.mapperFiles, mf)
		}
	}
	//parse mapper
	b.parseMappers()
	//prepare mapper
	b.prepareMappers()
	return b
}
