package gobatis

type txMapper struct {
	tx     *Tx
	mapper *mapper
}

// Update get update mapper
func (mapper *txMapper) Update(id string) *UpdateMapper {
	updateMapper := mapper.mapper.Update(id)
	if updateMapper != nil {
		updateMapper.tx = mapper.tx
	}
	return updateMapper
}

// Commit get update mapper
func (mapper *txMapper) Commit() error {
	return mapper.tx.commit()
}

// Rollback get update mapper
func (mapper *txMapper) Rollback() error {
	return mapper.tx.rollback()
}

// TxMapper get txMapper
func (b *Batis) TxMapper(binding string) *txMapper {
	m := b.Mapper(binding)
	db := m.currentDS.db
	tx := db.Begin().tx
	return &txMapper{
		tx: &Tx{
			db: db.db,
			tx: tx,
		},
		mapper: m,
	}
}
