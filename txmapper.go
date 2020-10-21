package gobatis

//Define txMapper struct
type txMapper struct {
	tx     *Tx
	mapper mapper
}

//Get update mapper
func (mapper *txMapper) Update(id string) *updateMapper {
	updateMapper := mapper.mapper.Update(id)
	if updateMapper != nil {
		updateMapper.tx = mapper.tx
	}
	return updateMapper
}

//Get update mapper
func (mapper *txMapper) Commit() error {
	return mapper.tx.commit()
}

//Get update mapper
func (mapper *txMapper) Rollback() error {
	return mapper.tx.rollback()
}

//Get txMapper
func (b *Batis) TxMapper(binding string) *txMapper {
	mapper := b.Mapper(binding)
	db := mapper.ds.db.db
	tx := mapper.ds.db.Begin().tx
	return &txMapper{
		tx: &Tx{
			db: db,
			tx: tx,
		},
		mapper: *mapper,
	}
}
