package gobatis

import (
	"fmt"
	"testing"
)

func TestDs(t *testing.T) {
	b := Default().Init().DSN(dsn)
	m := b.Mapper("user")
	time := ""
	m.Select("selectSimple").Exec().Scan(&time)
	fmt.Println(time)
}
