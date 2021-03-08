package gobatis

import (
	"fmt"
	"sync"
)

var mutex1 = &sync.Mutex{}
var mutex2 = &sync.Mutex{}

type mapper struct {
	logger        *log
	printSql      bool
	binding       string
	currentDSName string
	currentDS     *DS
	multiDS       *multiDS
	updateMappers map[string]*UpdateMapper
	selectMappers map[string]*SelectMapper
}

// Select get select mapper
func (m *mapper) Select(id string) *SelectMapper {
	return m.SelectWithDS(id, "")
}

// SelectWithDS select mapper with ds
func (m *mapper) SelectWithDS(id, ds string) *SelectMapper {
	defer mutex1.Unlock()
	mutex1.Lock()
	selectMapper, have := m.selectMappers[id]
	if !have {
		m.logger.Error("no select node : %v", id)
		panic(fmt.Sprintf("no select node : %v", id))
	}
	cloneSM := &SelectMapper{}
	copySelectMapper(cloneSM, selectMapper)
	if ds == "" {
		//set default db
		dsName, db := m.multiDS.defaultDS()
		m.logger.Info("[MultiDS]Choose DS[%s]", dsName)
		cloneSM.db = db.db
	} else {
		mds, mdsHave := m.multiDS.mds[ds]
		if !mdsHave {
			m.logger.Error("[MultiDS] DS[%s] was not registered", ds)
			panic(fmt.Sprintf("[MultiDS] DS[%s] was not registered", ds))
			return nil
		}
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		cloneSM.db = mds.db
	}
	cloneSM.printSql = m.printSql
	return cloneSM
}

// Update get update mapper
func (m *mapper) Update(id string) *UpdateMapper {
	return m.UpdateWithDS(id, "")
}

// UpdateWithDS get update mapper with ds
func (m *mapper) UpdateWithDS(id, ds string) *UpdateMapper {
	defer mutex2.Unlock()
	mutex2.Lock()
	updateMapper, have := m.updateMappers[id]
	if !have {
		m.logger.Error("no update node : %v", id)
		panic(fmt.Sprintf("no update node : %v", id))
	}
	cloneUM := &UpdateMapper{}
	copyUpdateMapper(cloneUM, updateMapper)
	if ds == "" {
		//set default db
		dsName, db := m.multiDS.defaultDS()
		m.logger.Info("[MultiDS]Choose DS[%s]", dsName)
		cloneUM.db = db.db
	} else {
		mds, mdsHave := m.multiDS.mds[ds]
		if !mdsHave {
			m.logger.Error("[MultiDS] DS[%s] was not registered", ds)
			panic(fmt.Sprintf("[MultiDS] DS[%s] was not registered", ds))
			return nil
		}
		m.logger.Info("[MultiDS]Choose DS[%s]", ds)
		cloneUM.db = mds.db
	}
	cloneUM.printSql = m.printSql
	return cloneUM
}

func copySelectMapper(dst, src *SelectMapper) {
	dst.db = src.db
	dst.logger = src.logger
	dst.binding = src.binding
	dst.extraSql = src.extraSql
	dst.gfuncMap = src.gfuncMap
	dst.id = src.id
	dst.originalSql = src.originalSql
	dst.sql = src.sql
	dst.printSql = src.printSql
}

func copyUpdateMapper(dst, src *UpdateMapper) {
	dst.db = src.db
	dst.logger = src.logger
	dst.binding = src.binding
	dst.gfuncMap = src.gfuncMap
	dst.id = src.id
	dst.originalSql = src.originalSql
	dst.sql = src.sql
	dst.printSql = src.printSql
}
