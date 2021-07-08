package gobatis

import (
	"github.com/sirupsen/logrus"
	"text/template"
	"time"
)

var batis *Batis

// Batis struct
type Batis struct {
	mappers     map[string]*mapper
	mapperNodes map[string]*mapperNode
	Logger      *logrus.Logger
	MultiDS     *multiDS
	FuncMap     template.FuncMap
}

func newBatis() *Batis {
	return &Batis{
		mappers:     make(map[string]*mapper, 0),
		mapperNodes: make(map[string]*mapperNode, 0),
		Logger:      logrus.StandardLogger(),
		MultiDS: &multiDS{
			mds:             make(map[string]*DS, 0),
			maxOpenConn:     10,
			maxIdleConn:     2,
			connMaxLifetime: 0,
			connMaxIdleTime: time.Minute * 2,
		},
		FuncMap: make(map[string]interface{}, 0),
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
		b.Logger.Panicf("mapper: no binding [%v]", binding)
	}
	_, mds := b.MultiDS.defaultDS()
	mp.currentDS = mds
	return mp
}
