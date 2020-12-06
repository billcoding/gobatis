package gobatis

import (
	"fmt"
	"testing"
)

func TestRaw(t *testing.T) {
	var xml1 = `<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="user">
    <select id="select">
        select now() as c
    </select>
</batis-mapper>`
	b := Default().AddRaw(xml1).DSN(dsn)
	b.Config.AutoScan = false
	b.Config.PrintSql = true
	b.Init()
	fmt.Println(b.Mapper("user").Select("select").Exec().SingleString())
}
