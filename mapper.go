package gobatis

type mapper struct {
	logger        *log
	printSql      bool
	binding       string
	currentDSName string
	currentDS     *DS
	multiDS       *multiDS
	updateMappers map[string]*updateMapper
	selectMappers map[string]*selectMapper
}

// Select get select mapper
func (m *mapper) Select(id string) *selectMapper {
	return m.SelectWithDS(id, "")
}

// SelectWithDS select mapper with ds
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
		mds, have := m.multiDS.mds[ds]
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

// Update get update mapper
func (m *mapper) Update(id string) *updateMapper {
	return m.UpdateWithDS(id, "")
}

// UpdateWithDS get update mapper with ds
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
		mds, have := m.multiDS.mds[ds]
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
