package gobatis

import "time"

func (b *Batis) SetMaxOpenConn(n int) *Batis {
	b.MultiDS.maxOpenConn = n
	return b
}

func (b *Batis) SetMaxIdleConn(n int) *Batis {
	b.MultiDS.maxIdleConn = n
	return b
}

func (b *Batis) SetConnMaxIdleTime(n time.Duration) *Batis {
	b.MultiDS.connMaxIdleTime = n
	return b
}

func (b *Batis) SetConnMaxLifetime(n time.Duration) *Batis {
	b.MultiDS.connMaxLifetime = n
	return b
}
