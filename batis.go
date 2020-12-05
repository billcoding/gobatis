package gobatis

import (
	l "log"
	"os"
	"sync"
)

var batis *Batis

//Define batis struct
type Batis struct {
	mutex             sync.Mutex             //mutex
	Config            *Config                //config
	Logger            *log                   //Logger
	MultiDS           MultiDS                //multiple datasource
	parsedMapperPaths []string               //parsed mapper path
	mapperFiles       []string               //mapperNode files
	mappers           map[string]*mapper     //mapper
	mapperNodes       map[string]*mapperNode //mapper nodes
}

//new batis
func newBatis() *Batis {
	return &Batis{
		mutex: sync.Mutex{},
		Config: &Config{
			PrintSql:    false,
			MapperPaths: []string{"./mapper"},
		},
		Logger: &log{
			ologger: l.New(os.Stdout, "[GOBATIS]", l.Flags()),
			elogger: l.New(os.Stdout, "[GOBATIS]", l.Flags()),
		},
		MultiDS:     make(map[string]*DS, 0),
		mappers:     make(map[string]*mapper, 0),
		mapperNodes: make(map[string]*mapperNode, 0),
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

//Start scan mapper file for binding
func (b *Batis) scanMapper() *Batis {
	if len(b.parsedMapperPaths) <= 0 {
		return b
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	//collect mapperNode files
	for _, mapperPath := range b.parsedMapperPaths {
		b.Logger.Info("[Mapper]scan mapper files : %v", mapperPath)
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

//Get mapper
func (b *Batis) Mapper(binding string) *mapper {
	mapper, have := b.mappers[binding]
	if !have {
		b.Logger.Error("[Mapper]no binding : %v", binding)
		return nil
	}

	_, mds := b.MultiDS.defaultDS()
	mapper.currentDS = mds
	mapper.printSql = b.Config.PrintSql
	return mapper
}

//Set mapper path
func (b *Batis) MapperPaths(mapperPaths ...string) *Batis {
	b.Config.MapperPaths = mapperPaths
	return b
}

//Parse mapper paths
func (b *Batis) parseMapperPaths() *Batis {
	for _, mapperPath := range b.Config.MapperPaths {
		b.parsedMapperPaths = append(b.parsedMapperPaths, mapperPath)
	}
	return b
}
