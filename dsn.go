package gobatis

//Register dsn with master ds
func (b *Batis) DSN(dsn string) *Batis {
	b.MultiDS.Add("master", dsn)
	return b
}

//Register dsn with master ds and dialect
func (b *Batis) DSNWithDialect(dialect Dialect, dsn string) *Batis {
	b.MultiDS.AddWithDialect("master", dsn, dialect)
	return b
}
