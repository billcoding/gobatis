package gobatis

import (
	"embed"
	"encoding/xml"
)

// AddFS add FS xml
func (b *Batis) AddFS(FS *embed.FS, name string) *Batis {
	node := mapperNode{}
	bytes, err := FS.ReadFile(name)
	if err != nil {
		b.Logger.Error("%v", err)
		panic(err)
	}
	err = xml.Unmarshal(bytes, &node)
	if err != nil {
		b.Logger.Error("%v", err)
		panic(err)
	}
	_, have := b.mappers[node.Binding]
	if have {
		b.Logger.Error("[FS]AddFS binding[%v] fail: duplicated", node.Binding)
		panic(err)
	}
	b.mappers[node.Binding] = &mapper{
		logger:        b.Logger,
		binding:       node.Binding,
		multiDS:       b.MultiDS,
		selectMappers: b.prepareSelectMappers(node.Binding, node.MapperSelectNodes),
		updateMappers: b.prepareUpdateMappers(node.Binding, node.MapperUpdateNodes),
	}
	b.Logger.Info("[FS]AddFS binding[%v] success", node.Binding)
	return b
}
