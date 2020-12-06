package gobatis

import (
	"fmt"
	"testing"
)

func TestHelper(t *testing.T) {
	var xml1 = `<?xml version="1.0" encoding="UTF-8"?>
<batis-mapper binding="user">
    <select id="select">
        select now() as c
    </select>
</batis-mapper>`
	batis := Default().DSN(dsn)
	batis.AddRaw(xml1)
	batis.Config.AutoScan = false
	batis.Init()
	var sm = NewHelper("user", "select").Select()
	fmt.Println(sm.Exec().SingleString())
}
