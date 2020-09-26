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
		batis.LogFatal("unregistered ds : %v", dsName)
		return nil
	}
	batis.LogInfo("using ds : %v", dsName)
	mapper.ds = *ds
	return mapper
}
