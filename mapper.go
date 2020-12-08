package gobatis

//Define mapper struct
type mapper struct {
	logger        *log                     //logger
	printSql      bool                     //print sql
	binding       string                   //binding name
	currentDSName string                   //current ds name
	currentDS     *DS                      //current ds
	multiDS       MultiDS                  //multi ds
	updateMappers map[string]*updateMapper //update mappers
	selectMappers map[string]*selectMapper //select mappers
}

//Get select mapper
func (m *mapper) Select(id string) *selectMapper {
	return m.SelectWithDS(id, "")
}

//Get select mapper with ds
func (m *mapper) SelectWithDS(id, ds string) *selectMapper {
	selectMapper, have := m.selectMappers[id]
	if !have {
		m.logger.Error("no select node : %v", id)
		return nil
	}
	if ds == "" {
		//set default db
		ds, db := m.multiDS.defaultDS()
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		selectMapper.db = db.db
	} else {
		mds, have := m.multiDS[ds]
		if !have {
			m.logger.Error("[MultiDS] DS[%s] was not registered", ds)
			return nil
		}
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		selectMapper.db = mds.db
	}
	selectMapper.printSql = m.printSql
	return selectMapper
}

//Get update mapper
func (m *mapper) Update(id string) *updateMapper {
	return m.UpdateWithDS(id, "")
}

//Get update mapper with ds
func (m *mapper) UpdateWithDS(id, ds string) *updateMapper {
	updateMapper, have := m.updateMappers[id]
	if !have {
		m.logger.Error("no update node : %v", id)
		return nil
	}
	if ds == "" {
		//set default db
		ds, db := m.multiDS.defaultDS()
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		updateMapper.db = db.db
	} else {
		mds, have := m.multiDS[ds]
		if !have {
			m.logger.Error("[MultiDS] DS[%s] was not registered", ds)
			return nil
		}
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		updateMapper.db = mds.db
	}
	updateMapper.printSql = m.printSql
	return updateMapper
}
