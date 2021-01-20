package gobatis

type TXMapper struct {
	tx     *TX
	mapper *mapper
}

// Update get update mapper
func (mapper *TXMapper) Update(id string) *UpdateMapper {
	updateMapper := mapper.mapper.Update(id)
	if updateMapper != nil {
		updateMapper.tx = mapper.tx
	}
	return updateMapper
}

// Commit get update mapper
func (mapper *TXMapper) Commit() error {
	return mapper.tx.commit()
}

// Rollback get update mapper
func (mapper *TXMapper) Rollback() error {
	return mapper.tx.rollback()
}

// TxMapper get TXMapper
func (b *Batis) TxMapper(binding string) *TXMapper {
	m := b.Mapper(binding)
	db := m.currentDS.db
	tx := db.Begin().tx
	return &TXMapper{
		tx: &TX{
			db: db.db,
			tx: tx,
		},
		mapper: m,
	}
}
