package gobatis

//Define Config struct
type Config struct {
	AutoScan    bool     //Auto scan mappers
	PrintSql    bool     //Print sql
	MapperPaths []string //Mapper Paths
}
