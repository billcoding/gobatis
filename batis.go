package gobatis

import (
	"github.com/sirupsen/logrus"
	"text/template"
)

var batis *Batis

// Batis struct
type Batis struct {
	mappers     map[string]*mapper
	mapperNodes map[string]*mapperNode
	Logger      *logrus.Logger
	MultiDS     *multiDS
	FuncMap     template.FuncMap
	PrintSql    bool
}

func newBatis() *Batis {
	return &Batis{
		mappers:     make(map[string]*mapper, 0),
		mapperNodes: make(map[string]*mapperNode, 0),
		Logger:      logrus.StandardLogger(),
		MultiDS: &multiDS{
			mds:    make(map[string]*DS, 0),
			config: &DBConfig{},
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
		b.Logger.Panicf("[Mapper]no binding : %v", binding)
	}
	_, mds := b.MultiDS.defaultDS()
	mp.currentDS = mds
	mp.printSql = b.PrintSql
	return mp
}
