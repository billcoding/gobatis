package gobatis

import (
	"os"
	"strings"
)

const (
	batisPrintSql = "BATIS_PRINT_SQL" //batis print sql
	batisDsn      = "BATIS_DSN"       //batis dsn
)

func (b *Batis) parseEnv() *Batis {
	printSql := os.Getenv(batisPrintSql)
	if printSql != "" {
		ss := strings.ToUpper(printSql)
		if ss == "ON" || ss == "TRUE" || ss == "1" {
			b.PrintSql = true
		}
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
