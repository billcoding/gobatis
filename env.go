package gobatis

import (
	"os"
	"strings"
)

const (
	batisPrintSql   = "BATIS_PRINT_SQL"   //batis print sql
	batisAutoScan   = "BATIS_AUTO_SCAN"   //batis auto scan
	batisMapperPath = "BATIS_MAPPER_PATH" //batis mapper path
	batisDsn        = "BATIS_DSN"         //batis dsn
)

//parse env variable for gobatis
func (b *Batis) parseEnv() *Batis {
	printSql := os.Getenv(batisPrintSql)
	if printSql != "" {
		ss := strings.ToUpper(printSql)
		if ss == "ON" || ss == "TRUE" || ss == "1" {
			b.Config.PrintSql = true
		}
	}
	autoScan := os.Getenv(batisAutoScan)
	if autoScan != "" {
		ss := strings.ToUpper(autoScan)
		if ss == "ON" || ss == "TRUE" || ss == "1" {
			b.Config.AutoScan = true
		}
	}
	mapperPath := os.Getenv(batisMapperPath)
	if mapperPath != "" {
		b.MapperPaths(strings.Split(mapperPath, ",")...)
	}
	batisDsn := os.Getenv(batisDsn)
	if batisDsn != "" {
		//NAME1,DSN1|NAME2,DSN2
		dsns := strings.Split(batisDsn, "|")
		if len(dsns) > 0 {
			if len(dsns) == 1 && !strings.Contains(dsns[0], ",") {
				//only one
				b.MultiDS.Add("master", dsns[0])
			} else {
				for _, dsnStr := range dsns {
					dsnArray := strings.Split(dsnStr, ",")
					name := "_"
					dsn := ""
					if len(dsnArray) > 1 {
						name = dsnArray[0]
						dsn = dsnArray[1]
					} else if len(dsnArray) == 1 {
						dsn = dsnArray[0]
					}
					b.MultiDS.Add(name, dsn)
				}
			}
		}
	}
	return b
}
