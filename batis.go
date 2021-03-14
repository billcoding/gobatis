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
	mutex       sync.Mutex
	mappers     map[string]*mapper
	mapperNodes map[string]*mapperNode
	Logger      *log
	MultiDS     *multiDS
	FuncMap     template.FuncMap
	PrintSql    bool
}

func newBatis() *Batis {
	return &Batis{
		mutex:       sync.Mutex{},
		mappers:     make(map[string]*mapper, 0),
		mapperNodes: make(map[string]*mapperNode, 0),
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
		FuncMap:  make(map[string]interface{}, 0),
		PrintSql: false,
	}
}

func init() {
	batis = newBatis()
}

// Default return default Batis
func Default() *Batis {
	if batis != nil {
		batis.parseEnv()
	}
	return batis
}

// New return new Batis
func New() *Batis {
	b := newBatis()
	batis.parseEnv()
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
	mp.printSql = b.PrintSql
	return mp
}
