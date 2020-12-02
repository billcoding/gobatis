package gobatis

//Define ds struct
type ds struct {
	dsn string //DSN for ds
	db  *DB    //sql db
}

//Select ds
func (mapper *mapper) SelectDS(dsName string) *mapper {
	ds, have := batis.dss[dsName]
	if !have {
		batis.Error("unregistered ds : %v", dsName)
		return nil
	}
	batis.Info("using ds : %v", dsName)
	mapper.ds = *ds
	return mapper
}
