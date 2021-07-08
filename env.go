package gobatis

import (
	"os"
	"strings"
)

const (
	envBatisDSN = "BATIS_DSN"
)

func (b *Batis) parseEnv() *Batis {
	if batisDSN := os.Getenv(envBatisDSN); batisDSN != "" {
		//NAME1,DSN1|NAME2,DSN2
		dsnS := strings.Split(batisDSN, "|")
		if len(dsnS) > 0 {
			if len(dsnS) == 1 && !strings.Contains(dsnS[0], ",") {
				//only one
				b.MultiDS.Add("master", dsnS[0])
			} else {
				for _, dsnStr := range dsnS {
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
