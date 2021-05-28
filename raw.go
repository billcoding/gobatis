package gobatis

import "encoding/xml"

// AddRaw add raw xml
func (b *Batis) AddRaw(rawXML string) *Batis {
	node := mapperNode{}
	err := xml.Unmarshal([]byte(rawXML), &node)
	if err != nil {
		b.Logger.Panicf("%v", err)
	}
	_, have := b.mappers[node.Binding]
	if have {
		b.Logger.Panicf("[Raw]AddRaw binding[%v] fail: duplicated", node.Binding)
	}
	b.mapperNodes[node.Binding] = &node
	b.mappers[node.Binding] = &mapper{
		logger:        b.Logger,
		binding:       node.Binding,
		multiDS:       b.MultiDS,
		selectMappers: b.prepareSelectMappers(node.Binding, node.MapperSelectNodes),
		updateMappers: b.prepareUpdateMappers(node.Binding, node.MapperUpdateNodes),
	}
	b.Logger.Infof("[Raw]AddRaw binding[%v] success", node.Binding)
	return b
}
