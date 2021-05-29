package gobatis

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

func (b *Batis) prepareUpdateMappers(binding string, mapperUpdateNodes []mapperUpdateNode) map[string]*UpdateMapper {
	updateMapperMap := make(map[string]*UpdateMapper, 0)
	if mapperUpdateNodes != nil {
		for _, node := range mapperUpdateNodes {
			id := node.Id
			sql := node.Text
			if sql == "" {
				b.Logger.Warnf("[Mapper]node sql is empty : %v", id)
				continue
			}
			updateMapperMap[id] = &UpdateMapper{
				funcMap:     &b.FuncMap,
				printSql:    b.PrintSql,
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

func (b *Batis) prepareSelectMappers(binding string, mapperSelectNodes []mapperSelectNode) map[string]*SelectMapper {
	selectMapperMap := make(map[string]*SelectMapper, 0)
	if mapperSelectNodes != nil {
		for _, node := range mapperSelectNodes {
			id := node.Id
			sql := node.Text
			if sql == "" {
				b.Logger.Warnf("[Mapper]node sql is empty : %v", id)
				continue
			}
			selectMapperMap[id] = &SelectMapper{
				funcMap:     &b.FuncMap,
				logger:      b.Logger,
				printSql:    b.PrintSql,
				binding:     binding,
				id:          id,
				originalSql: sql,
				sql:         sql,
			}
		}
	}
	return selectMapperMap
}
