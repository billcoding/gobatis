package gobatis

import (
	l "log"
	"os"
	"sync"
	"text/template"
)

var batis *Batis

// Batis struct
type Batis struct {
	inited            bool
	mutex             sync.Mutex
	parsedMapperPaths []string
	mapperFiles       []string
	mappers           map[string]*mapper
	mapperNodes       map[string]*mapperNode
	// Config field
	Config *Config
	// Logger field
	Logger *log
	// MultiDS field
	MultiDS *multiDS
	// FuncMap field
	FuncMap template.FuncMap
}

func newBatis() *Batis {
	return &Batis{
		mutex:             sync.Mutex{},
		parsedMapperPaths: make([]string, 0),
		mapperFiles:       make([]string, 0),
		mappers:           make(map[string]*mapper, 0),
		mapperNodes:       make(map[string]*mapperNode, 0),
		Config: &Config{
			AutoScan:    true,
			PrintSql:    false,
			MapperPaths: []string{"./mapper"},
		},
		Logger: &log{
			outLogger: l.New(os.Stdout, "[GOBATIS]", l.LstdFlags),
			errLogger: l.New(os.Stdout, "[GOBATIS]", l.LstdFlags),
		},
		MultiDS: &multiDS{
			mds: make(map[string]*DS, 0),
			config: &DBConfig{
				MaxIdleConns:    2,
				MaxOpenConns:    10,
				ConnMaxLifetime: 10,
			},
		},
		FuncMap: make(map[string]interface{}, 0),
	}
}

func init() {
	batis = newBatis()
}

// Default return default Batis
func Default() *Batis {
	return batis
}

// New return new Batis
func New() *Batis {
	return newBatis()
}

// Init Batis
func (b *Batis) Init() *Batis {
	b.parseEnv()
	b.parseMapperPaths()
	b.scanMapper()
	b.inited = true
	return b
}

func (b *Batis) scanMapper() *Batis {
	if !b.Config.AutoScan {
		return b
	}
	if len(b.parsedMapperPaths) <= 0 {
		return b
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	// collect mapperNode files
	for _, mapperPath := range b.parsedMapperPaths {
		b.Logger.Info("[Mapper]scan mapper files : %v", mapperPath)
		for _, mf := range getMapperFiles(mapperPath) {
			b.mapperFiles = append(b.mapperFiles, mf)
		}
	}
	b.parseMappers()
	b.prepareMappers()
	return b
}

// Mapper return cached mapper
func (b *Batis) Mapper(binding string) *mapper {
	mp, have := b.mappers[binding]
	if !have {
		b.Logger.Error("[Mapper]no binding : %v", binding)
		return nil
	}
	_, mds := b.MultiDS.defaultDS()
	mp.currentDS = mds
	mp.printSql = b.Config.PrintSql
	return mp
}

// MapperPaths set mapper path
func (b *Batis) MapperPaths(mapperPaths ...string) *Batis {
	b.Config.MapperPaths = mapperPaths
	return b
}

func (b *Batis) parseMapperPaths() *Batis {
	for _, mapperPath := range b.Config.MapperPaths {
		b.parsedMapperPaths = append(b.parsedMapperPaths, mapperPath)
	}
	return b
}
