package gobatis

type Dialect string

const (
	MySQL   Dialect = "mysql"     //see  https://github.com/go-sql-driver/mysql
	SQLite  Dialect = "sqlite3"   //see  https://github.com/mattn/go-sqlite3
	SQLite3 Dialect = "sqlite3"   //see  https://github.com/mattn/go-sqlite3
	MSSQL   Dialect = "sqlserver" //see  https://github.com/denisenkom/go-mssqldb
)
