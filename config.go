package gobatis

import "time"

// Config struct
type Config struct {
	// AutoScan field
	AutoScan bool
	// PrintSql field
	PrintSql bool
	// MapperPaths field
	MapperPaths []string
}

// DBConfig struct
type DBConfig struct {
	// MaxOpenConns field
	MaxOpenConns int
	// MaxIdleConns field
	MaxIdleConns int
	// ConnMaxLifetime field
	ConnMaxLifetime time.Duration
	// ConnMaxIdleTime field
	ConnMaxIdleTime time.Duration
}

// DBConfig config
func (b *Batis) DBConfig(config *DBConfig) *Batis {
	b.MultiDS.config = config
	return b
}
