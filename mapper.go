package gobatis

//Define mapper struct
type mapper struct {
	binding       string                  //binding name
	ds            ds                      //mapper ds
	updateMappers map[string]updateMapper //update mappers
	selectMappers map[string]selectMapper //select mappers
}

//Get select mapper
func (mapper *mapper) Select(id string) *selectMapper {
	selectMapper, have := mapper.selectMappers[id]
	if !have {
		batis.LogFatal("no select node : %v", id)
		return nil
	}
	//set db
	selectMapper.db = mapper.ds.db
	return &selectMapper
}

//Get update mapper
func (mapper *mapper) Update(id string) *updateMapper {
	updateMapper, have := mapper.updateMappers[id]
	if !have {
		batis.LogFatal("no update node : %v", id)
		return nil
	}
	//set db
	updateMapper.db = mapper.ds.db
	return &updateMapper
}

//Get select mapper
func (mapper *txMapper) Select(id string) *selectMapper {
	selectMapper := mapper.mapper.Select(id)
	if selectMapper != nil {
		selectMapper.tx = mapper.tx
	}
	return selectMapper
}

//Get mapper
func (b *Batis) Mapper(binding string) *mapper {
	mapper, have := b.mappers[binding]
	if !have {
		b.LogFatal("no binding : %v", binding)
		return nil
	}

	//init ds
	if len(b.dss) == 1 {
		for _, ds := range b.dss {
			mapper.ds = *ds
			break
		}
	} else if len(b.dss) > 1 {
		for name, ds := range b.dss {
			if name == "master" {
				mapper.ds = *ds
				break
			}
		}
	}
	return &mapper
}

//Set mapper path
func (b *Batis) MapperPaths(mapperPaths ...string) *Batis {
	b.mapperPaths = mapperPaths
	return b
}

//Parse mapper paths
func (b *Batis) parseMapperPaths() *Batis {
	for _, mapperPath := range b.mapperPaths {
		b.parsedMapperPaths = append(b.parsedMapperPaths, mapperPath)
	}
	return b
}
