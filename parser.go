package gobatis

import (
	"encoding/xml"
	"io/ioutil"
)

//parse mapper file with xml
func (b *Batis) parseMappers() *Batis {
	for _, file := range b.mapperFiles {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			b.Logger.Error("[Mapper]error : %v", err)
			continue
		}
		mapperNode := &mapperNode{}
		b.Logger.Info("[Mapper]parsing mapper file : %v", file)
		err = xml.Unmarshal(bytes, &mapperNode)
		if err != nil {
			b.Logger.Error("[Mapper]error : %v", err)
			continue
		}
		if mapperNode.Binding == "" {
			b.Logger.Error("[Mapper]mapper binding muse be provided in `%v`", file)
			continue
		}
		if _, have := b.mapperNodes[mapperNode.Binding]; have {
			b.Logger.Error("[Mapper]mapper binding is exists in `%v`", file)
			continue
		}
		b.mapperNodes[mapperNode.Binding] = mapperNode
	}
	return b
}

//Prepare mapper
func (b *Batis) prepareMappers() {
	for binding, node := range b.mapperNodes {
		updateMappers := b.prepareUpdateMappers(binding, node.MapperUpdateNodes)
		selectMappers := b.prepareSelectMappers(binding, node.MapperSelectNodes)
		if len(updateMappers) <= 0 && len(selectMappers) <= 0 {
			continue
		}
		b.mappers[binding] = &mapper{
			logger:        b.Logger,
			binding:       binding,
			multiDS:       b.MultiDS,
			updateMappers: updateMappers,
			selectMappers: selectMappers,
		}
	}
}

func (b *Batis) prepareUpdateMappers(binding string, mapperUpdateNodes []mapperUpdateNode) map[string]*updateMapper {
	updateMapperMap := make(map[string]*updateMapper, 0)
	if mapperUpdateNodes != nil {
		for _, node := range mapperUpdateNodes {
			id := node.Id
			sql := node.Text
			if sql == "" {
				b.Logger.Warn("[Mapper]node sql is empty : %v", id)
				continue
			}
			updateMapperMap[id] = &updateMapper{
				gfuncMap:    b.FuncMap,
				printSql:    b.Config.PrintSql,
				logger:      b.Logger,
				binding:     binding,
				id:          id,
				originalSql: sql,
				sql:         sql,
			}
		}
	}
	return updateMapperMap
}

func (b *Batis) prepareSelectMappers(binding string, mapperSelectNodes []mapperSelectNode) map[string]*selectMapper {
	selectMapperMap := make(map[string]*selectMapper, 0)
	if mapperSelectNodes != nil {
		for _, node := range mapperSelectNodes {
			id := node.Id
			sql := node.Text
			if sql == "" {
				b.Logger.Warn("[Mapper]node sql is empty : %v", id)
				continue
			}
			selectMapperMap[id] = &selectMapper{
				gfuncMap:    b.FuncMap,
				logger:      b.Logger,
				printSql:    b.Config.PrintSql,
				binding:     binding,
				id:          id,
				originalSql: sql,
				sql:         sql,
			}
		}
	}
	return selectMapperMap
}
